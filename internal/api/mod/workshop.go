package mod

import (
	"github.com/dstgo/steamapi"
	"github.com/dstgo/tracker/internal/types/dst"
)

type WorkShopModsLogic struct {
	steamClient *steamapi.Client
}

func (w WorkShopModsLogic) Search(searchOption dst.SearchModsOption) (dst.SearchModsResult, error) {
	//TODO implement me
	panic("implement me")
}

func (w WorkShopModsLogic) Details() (dst.ModsDetail, error) {
	//TODO implement me
	panic("implement me")
}

func (w WorkShopModsLogic) Options() (any, error) {
	//TODO implement me
	panic("implement me")
}

func (w WorkShopModsLogic) Version() string {
	//TODO implement me
	panic("implement me")
}
