package repo

import (
	"context"
	"github.com/dstgo/tracker/internal/types"
	"github.com/dstgo/tracker/pkg/lobbyapi"
	"github.com/qiniu/qmgo"
	opts "github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type LobbyServer struct {
	// geo info
	Region       string   `bson:"region"`
	Continent    string   `bson:"continent"`
	Area         string   `bson:"area"`
	City         string   `bson:"city"`
	PlatformName string   `bson:"platform_name"`
	TagNames     []string `bson:"tag_names"`

	// created at timestamp
	CreatedAt       int64 `bson:"created_at"`
	lobbyapi.Server `bson:"inline"`
}

type LobbyServerDetails struct {
	Area                   string
	CreatedAt              int64
	lobbyapi.ServerDetails `bson:"inline"`
}

// NewLobbyRepo returns new lobby mongo db operator
func NewLobbyRepo(ctx context.Context, db *qmgo.QmgoClient) (*LobbyRepo, error) {
	col := db.Database.Collection("lobby")

	// create index
	err := col.CreateIndexes(ctx, []opts.IndexModel{
		{[]string{"name"}, &options.IndexOptions{}},
		{[]string{"area"}, &options.IndexOptions{}},
		{[]string{"platform_name"}, &options.IndexOptions{}},
		{[]string{"tag_names"}, &options.IndexOptions{}},
		{[]string{"created_at"}, &options.IndexOptions{}},
		{[]string{"row_id"}, &options.IndexOptions{}},
		{[]string{"game_mode"}, &options.IndexOptions{}},
		{[]string{"intent"}, &options.IndexOptions{}},
	})

	if err != nil {
		return nil, err
	}

	return &LobbyRepo{cli: db, collection: col}, nil
}

type LobbyRepo struct {
	cli        *qmgo.QmgoClient
	collection *qmgo.Collection
}

// RemoveServers returns deletedCount and total count after removing the specified servers
func (l *LobbyRepo) RemoveServers(ctx context.Context, filter bson.M) (int64, int64, error) {
	result, err := l.collection.RemoveAll(ctx, filter)
	if err != nil {
		return 0, 0, err
	}

	estimatedCount, err := l.collection.Find(ctx, bson.M{}).EstimatedCount()
	if err != nil {
		return 0, 0, err
	}
	return result.DeletedCount, estimatedCount, nil
}

func (l *LobbyRepo) InsertManyServers(ctx context.Context, servers []LobbyServer) (int, error) {
	// do transaction
	result, err := l.collection.InsertMany(ctx, servers)
	if err != nil {
		return 0, err
	}
	return len(result.InsertedIDs), nil
}

// FindServers returns list of servers by page
func (l *LobbyRepo) FindServers(ctx context.Context, page, size int, sort string, filter bson.M) (types.PageResult[LobbyServer], error) {
	if page <= 0 {
		page = 1
	}

	if size <= 0 {
		size = 10
	}

	if sort == "" {
		sort = "name"
	}

	var result types.PageResult[LobbyServer]
	lastTs := bson.M{}

	// get the latest inserted timestamp
	err := l.collection.Aggregate(ctx, qmgo.Pipeline{
		bson.D{
			{"$group", bson.M{"_id": "$created_at"}},
		},
		bson.D{
			{"$sort", bson.M{"created": 1}},
		},
	}).One(&lastTs)

	if err != nil {
		return result, err
	}

	// mean to there has no data in database
	if len(lastTs) == 0 {
		return result, nil
	}

	// specify latest timestamp
	ts := lastTs["_id"]
	filter["created_at"] = ts

	// total count
	total, err := l.collection.Find(ctx, bson.M{"created_at": ts}).EstimatedCount()
	if err != nil {
		return result, err
	}
	result.Total = total

	// match
	matchStage := bson.D{{"$match", filter}}
	// distinct by grow_id and returns object_id for per item
	groupStage := bson.D{{"$group", bson.M{"_id": "$row_id", "object_id": bson.M{"$first": "$_id"}}}}
	// pagination
	skipStage := bson.D{{"$skip", (page - 1) * size}}
	limitStage := bson.D{{"$limit", size}}

	// filter results and distinct by row_id, then pagination
	var objs []bson.M
	err = l.collection.Aggregate(ctx, qmgo.Pipeline{matchStage, groupStage, skipStage, limitStage}).All(&objs)
	if err != nil {
		return result, err
	}

	// collect object_id
	var ids []any
	for _, obj := range objs {
		ids = append(ids, obj["object_id"])
	}

	// find final result
	err = l.collection.Find(ctx, bson.M{"_id": bson.M{"$in": ids}}).All(&result.List)
	if err != nil {
		return result, err
	}

	return result, nil
}

type LobbyStatisticItem struct {
	Label         string `json:"label:" bson:"label"`
	TotalServers  int64  `json:"totalServers" bson:"totalServers"`
	OnlinePlayers int64  `json:"onlinePlayers" bson:"onlinePlayers"`
}

type LobbyStatisticInfo struct {
	TotalServers  int64 `json:"totalServers" bson:"totalServers"`
	OnlinePlayers int64 `json:"onlinePlayers" bson:"onlinePlayers"`

	Platforms []LobbyStatisticItem `json:"platforms" bson:"platforms"`
	Area      []LobbyStatisticItem `json:"area" bson:"area"`
	Ts        int64                `json:"ts" bson:"ts"`
}

func NewLobbyStatisticRepo(cli *qmgo.QmgoClient) *LobbyStatisticRepo {
	return &LobbyStatisticRepo{col: cli.Database.Collection("lobby_sum")}
}

type LobbyStatisticRepo struct {
	col *qmgo.Collection
}

func (l *LobbyStatisticRepo) InsertOne(ctx context.Context, data LobbyStatisticInfo) error {
	_, err := l.col.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (l *LobbyStatisticRepo) GetMany(ctx context.Context, before, until, tail int64, duration time.Duration) ([]LobbyStatisticInfo, error) {
	var result []LobbyStatisticInfo

	if tail == 0 {
		tail = 100
	}

	// convert unit
	duration /= time.Millisecond

	err := l.col.Aggregate(ctx, qmgo.Pipeline{
		bson.D{
			{"$match", bson.M{"ts": bson.M{"$gte": before, "$lte": until, "$mod": bson.A{duration, 0}}}},
		},
		bson.D{
			{"$sort", bson.M{"ts": -1}},
		},
		bson.D{
			{"$limit", tail},
		},
	}).All(&result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
