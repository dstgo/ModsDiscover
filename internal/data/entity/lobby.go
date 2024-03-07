package entity

type LobbyServer struct {
	RowId string `gorm:"primaryKey;"`
	// created timestamp
	CreatedAt int64 `gorm:"primaryKey;"`

	// network options
	Guid string
	// only at steam platform
	SteamId string
	// only for clan server
	SteamClanId string
	// only for no password server
	OwnerNetId string
	SteamRoom  string
	Session    string
	Address    string
	Port       int
	Host       string
	Platform   int

	ClanOnly bool
	LanOnly  bool

	// game options
	Version  int
	Name     string
	GameMode string
	Intent   string
	Season   string
	Tags     string
	// max players allowed
	MaxConnections int
	// online players number
	Connected int

	Mod             bool
	Pvp             bool
	HasPassword     bool
	IsDedicated     bool
	ClientHosted    bool
	AllowNewPlayers bool
	ServerPaused    bool
	FriendOnly      bool
}
