package lobbyapi

const GameId = "DontStarveTogether"

// ExplicitPlatforms means that platforms could be used as query params in klei api.
var ExplicitPlatforms = []string{Steam.String(), PSN.String(), Rail.String(), XBOne.String(), Switch.String()}

const (
	Steam Platform = 1
	PSN   Platform = 2
	// Rail is alias of WeGame, only serve at ap-east-1
	Rail  Platform = 4
	XBOne Platform = 16
	// PS4Official can not be use in api query params
	PS4Official Platform = 19
	Switch      Platform = 32
)

// Platform represents dst server platform, it may be updated by klei in the future
type Platform uint

func (p Platform) String() string {
	switch p {
	case 1:
		return "Steam"
	case 2:
		return "PSN"
	case 4:
		return "Rail"
	case 16:
		return "XBone"
	case 19:
		return "PS4Official"
	case 32:
		return "Switch"
	}
	return "unknown platform"
}

// Region represents dst lobby server region, it may be updated by klei in the future
const (
	UsEast1     = "us-east-1"
	EuCentral   = "eu-central-1"
	ApSoutheast = "ap-southeast-1"
	ApEast      = "ap-east-1"
)

// Lobby server urls, may be updated by klei in the future
const (
	LobbyRegionURL  = `https://lobby-v2-cdn.klei.com/regioncapabilities-v2.json`
	LobbyServersURL = `https://lobby-v2-cdn.klei.com/{{.region}}-{{.platform}}.json.gz`
	LobbyDetailsURL = "https://lobby-v2-{{.region}}.klei.com/lobby/read"
)

type Regions struct {
	Regions []struct {
		Region string `json:"Region"`
	} `json:"LobbyRegions"`
}

// Server includes all the information about single dst server
type Server struct {
	// network options
	Guid  string `json:"guid" bson:"guid"`
	RowId string `json:"__rowId" bson:"row_id"`
	// only at steam platform
	SteamId string `json:"steamid" bson:"steam_id"`
	// only for clan server
	SteamClanId string `json:"steamclanid" bson:"steam_clan_id"`
	// only for no password server
	OwnerNetId string `json:"ownernetid" bson:"owner_net_id"`
	SteamRoom  string `json:"steamroom" bson:"steam_room"`
	Session    string `json:"session" bson:"session"`
	Address    string `json:"__addr" bson:"address"`
	Port       int    `json:"port" bson:"port"`
	Host       string `json:"host" bson:"host"`

	Platform Platform `json:"platform" bson:"platform"`

	ClanOnly bool `json:"clanonly" bson:"clan_only"`
	LanOnly  bool `json:"lanonly" bson:"lan_only"`

	// second shard
	Secondaries map[string]Secondaries `bson:"secondaries"`

	// game options
	Name     string `json:"name" bson:"name"`
	GameMode string `json:"mode" bson:"game_mode"`
	Intent   string `json:"intent" bson:"intent"`
	Season   string `json:"season" bson:"season"`
	Tags     string `json:"tags" bson:"-"`
	Version  int    `json:"v" bson:"version"`
	// max players allowed
	MaxConnections int `json:"maxconnections" bson:"max_connections"`
	// online players number
	Connected int `json:"connected" bson:"connected"`

	ModEnabled      bool `json:"mods" bson:"mod_enabled"`
	PvpEnabled      bool `json:"pvp" bson:"pvp_enabled"`
	HasPassword     bool `json:"password" bson:"has_password"`
	IsDedicated     bool `json:"dedicated" bson:"is_dedicated"`
	ClientHosted    bool `json:"clienthosted" bson:"client_hosted"`
	AllowNewPlayers bool `json:"allownewplayers" bson:"allow_new_players"`
	ServerPaused    bool `json:"serverpaused" bson:"server_paused"`
	FriendOnly      bool `json:"fo" bson:"friend_only"`
}

// Secondaries represents the secondaries shard among dst servers
type Secondaries struct {
	Id      string `json:"id" bson:"id"`
	SteamId string `json:"steamid" bson:"steam_id"`
	Address string `json:"__addr" bson:"address"`
	Port    int    `json:"port" bson:"port"`
}

type Servers struct {
	List []Server `json:"GET"`
}

type Player struct {
	Name    string `bson:"name" json:"name"`
	Prefab  string `bson:"prefab" json:"prefab"`
	SteamId string `bson:"steam_id" json:"steamId"`
	// hex color code
	Colour string `bson:"colour" json:"colour"`
	// shard level
	Level int `bson:"level" json:"level"`
}

type Mod struct {
	Id       string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Version1 string `bson:"version1" json:"version1"`
	Version2 string `bson:"version2" json:"version2"`
	Enabled  bool   `bson:"enabled" json:"enabled"`
}

type Details struct {
	Day                int      `bson:"day" json:"day"`
	DayElapsedInSeason int      `bson:"day_elapsed_in_season" json:"dayElapsedInSeason"`
	DaysLeftInSeason   int      `bson:"days_left_in_season" json:"daysLeftInSeason"`
	Players            []Player `bson:"players" json:"playerList"`
	Mods               []Mod    `bson:"mods" json:"modList"`
}

// ServerDetails includes some details information
type ServerDetails struct {
	// repeat options
	Server

	Tick          int  `json:"tick" bson:"tick"`
	ClientModsOff bool `json:"clientmodsoff" bson:"client_mods_off"`
	Nat           int  `json:"nat" bson:"nat"`

	// raw lua script data
	Data          string `json:"data" bson:"data"`
	WorldGen      string `json:"worldgen" bson:"world_gen"`
	OnlinePlayers string `json:"players" bson:"online_players"`
	Mods          []any  `json:"mods_info" bson:"mods"`

	// parsed lua data
	Details Details `bson:"details"`
}
