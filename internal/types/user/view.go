package user

import "github.com/dstgo/tracker/internal/types/role"

var InitialUser = CreateUserOption{
	Username: "dstadmin",
	Email:    "",
	Password: "0123456789",
}

type Info struct {
	UUID      string          `json:"uuid" example:"55BBA4ED-18D3-790F-EABF-A5330E527586"`
	Username  string          `json:"username" example:"jack"`
	Email     string          `json:"email" example:"jacklove@lol.com"`
	CreatedAt uint64          `json:"createdAt" example:"947416200"`
	Roles     []role.RoleInfo `json:"roles"`
}
