package types

import "github.com/dstgo/tracker/pkg/lobbyapi"

type QueryLobbyServersOptions struct {
	Page int    `query:"page" default:"1" binding:"gt=0"`
	Size int    `query:"size" default:"10" binding:"gt=0,lte=100"`
	Sort string `query:"sort" default:"name"`

	// network query options
	Address string `query:"address"`
	// area code
	Area string `query:"area"`
	// 0.All
	// 1-Steam
	// 2-WeGame
	// 3-PSN
	// 4-Xbox
	// 5-Ps4Official
	// 6-NS
	Platform int `query:"platform"`
	// 0 - all
	// 1 - dedicated
	// 2 - clienthosted
	// 3 - steamgroup
	ServerType int `query:"server_type"`

	// game query options
	Name string `query:"name"`
	// format like tag1,tag2,tag3,tag4,tag5
	Tags     string `query:"tags"`
	GameMode string `query:"game_mode"`
	Intent   string `query:"intent"`

	// -1 off
	//  0 ignored
	//  1 on
	PvpEnabled  int `query:"pvp"`
	ModEnabled  int `query:"mod"`
	HasPassword int `query:"password"`
}

type QueryLobbyServersResp struct {
	// network
	RowId       string `json:"rowId"`
	SteamClanId string `json:"steamClanId"`
	Address     string `json:"address"`
	Port        int    `json:"port"`
	Host        string `json:"host"`

	// geo information
	Region       string `json:"region"`
	Continent    string `json:"continent"`
	Area         string `json:"area"`
	City         string `json:"city"`
	PlatformName string `json:"PlatformName"`
	Platform     int    `json:"platform"`

	// game options
	Version    int      `json:"version"`
	Name       string   `json:"name"`
	GameMode   string   `json:"mode"`
	Intent     string   `json:"intent"`
	Season     string   `json:"season"`
	Tags       []string `json:"tags"`
	MaxPlayers int      `json:"maxPlayers"`
	Online     int      `json:"online"`

	// other properties
	Mod             bool `json:"mods"`
	Pvp             bool `json:"pvp"`
	HasPassword     bool `json:"password"`
	IsDedicated     bool `json:"dedicated"`
	ClientHosted    bool `json:"clientHosted"`
	AllowNewPlayers bool `json:"allowNewPlayers"`
	ServerPaused    bool `json:"serverPaused"`
	FriendOnly      bool `json:"friendOnly"`
	ClanOnly        bool `json:"clanOnly"`
}

type QueryLobbyServerDetailsOption struct {
	RowId  string `query:"rowId" binding:"required"`
	Region string `query:"region" binding:"required"`
}

type QueryLobbyServerDetailResp struct {
	QueryLobbyServersResp
	lobbyapi.Details
}

type QueryLobbyStatisticOption struct {
	Until  int64 `query:"until" binding:"gt=0"`
	Before int64 `query:"before" binding:"gt=0"`
	Tail   int64 `query:"tail" binding:"gt=0"`
}
