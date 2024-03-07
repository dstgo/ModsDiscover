package system

import (
	"github.com/dstgo/tracker/internal/types/helper"
	"github.com/dstgo/tracker/pkg/vax"
	"github.com/spf13/cast"
)

const (
	AppAPI = "appapi"

	OpenAPI = "openapi"
)

type PingRequest struct {
	// name must be one of [wilson, wendy]
	Name string `json:"name" uri:"name" form:"name" label:"field.name" example:"wilson"`
}

func (p PingRequest) Validate(lang string) error {
	return vax.Struct(&p, lang,
		vax.Field(&p.Name, helper.RequiredRules(RulePing)...),
	)
}

// Id
// represent query or path param ID
type Id struct {
	Id string `json:"id" uri:"id" form:"id" label:"field.id"`
}

func (i Id) Int() int {
	return cast.ToInt(i.Id)
}

func (i Id) Uint() uint {
	return cast.ToUint(i.Id)
}

func (i Id) String() string {
	return cast.ToString(i.Id)
}

func (i Id) Validate(lang string) error {
	return vax.Struct(&i, lang,
		vax.Field(&i.Id, vax.Required),
	)
}

type Uid struct {
	UUID string `json:"uuid" uri:"uuid" form:"uuid" label:"field.uuid"`
}

func (u Uid) Validate(lang string) error {
	return vax.Struct(&u, lang,
		vax.Field(&u.UUID, vax.Required),
	)
}
