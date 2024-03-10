package dst

type QueryLobbyServersOptions struct {
	Page int `uri:"page" form:"page"`
	Size int `uri:"size" form:"size"`

	// network query options
	Address  string `uri:"ip_address" form:"ip_address"`
	Country  string `uri:"country" form:"country"`
	Platform int    `uri:"platform" form:"platform"`
	// 0 - dedicated
	// 2 - clienthosted
	// 4 - steamclanid
	ServerType int `uri:"server_type" form:"server_type"`

	// game query options
	Name string `uri:"name" form:"name"`
	//
	Tags     string `uri:"tags" form:"tags"`
	GameMode string `uri:"game_mode" form:"game_mode"`
	Intent   string `uri:"intent" form:"intent"`
	Pvp      bool   `uri:"pvp" form:"pvp"`
	Mod      bool   `uri:"mod" form:"mod"`
	Password bool   `uri:"password" form:"password"`
}

type QueryLobbyServersResp struct {
	// network
	RowId       string `json:"rowId"`
	SteamClanId string `json:"steamClanId"`
	Address     string `json:"address"`
	Port        int    `json:"port"`
	Host        string `json:"host"`
	Platform    int    `json:"platform"`

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
