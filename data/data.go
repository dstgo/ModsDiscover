package data

import (
	"context"
	"fmt"
	"github.com/dstgo/tracker/conf"
	"github.com/go-redis/redis/v8"
	"github.com/qiniu/qmgo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadGormDB(dbConf conf.DBConf) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbConf.User, dbConf.Password, dbConf.Address, dbConf.Params)
	dialector := mysql.Open(dsn)
	db, err := gorm.Open(dialector)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func LoadMongoDB(ctx context.Context, dbConf conf.DBConf) (*qmgo.Client, error) {
	//  [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s?%s", dbConf.User, dbConf.Password, dbConf.Address, dbConf.DataBase, dbConf.Params)
	mgocli, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: uri})
	if err != nil {
		return nil, err
	}
	return mgocli, nil
}

func LoadRedisDB(ctx context.Context, rdConf conf.RedisConf) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     rdConf.Address,
		Password: rdConf.Auth,
	})
	pingRes := client.Ping(ctx)
	if err := pingRes.Err(); err != nil {
		return nil, err
	}
	return client, nil
}
