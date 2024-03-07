package user

import (
	"github.com/dstgo/tracker/internal/data"
)

func NewInfoLogic(ds *data.DataSource) InfoLogic {
	return InfoLogic{ds: ds}
}

type InfoLogic struct {
	ds *data.DataSource
}
