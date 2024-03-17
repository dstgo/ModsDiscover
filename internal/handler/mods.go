package handler

import (
	"context"
	"github.com/dstgo/steamapi"
	"github.com/dstgo/steamapi/types/publishedfile"
	"github.com/dstgo/steamapi/types/steam"
	"github.com/dstgo/tracker/internal/types"
)

type ModHandler interface {
	SearchModList(ctx context.Context, queryOption types.SearchModsOption) (publishedfile.FileList, error)
}

func NewWorkShopHandler(steamCLI *steamapi.Client) *WorkShopModHandler {
	return &WorkShopModHandler{steamCLI: steamCLI}
}

var _ ModHandler = (*WorkShopModHandler)(nil)

type WorkShopModHandler struct {
	steamCLI *steamapi.Client
}

func (w *WorkShopModHandler) SearchModList(ctx context.Context, queryOption types.SearchModsOption) (publishedfile.FileList, error) {
	// dst app id

	options := publishedfile.FileQueryOption{
		AppID:                  types.DstAppID,
		Language:               steam.LanguageCode(queryOption.Lang),
		SearchText:             queryOption.Text,
		NumPerPage:             uint(queryOption.Size),
		Page:                   uint(queryOption.Page),
		ReturnTags:             queryOption.Tags,
		ReturnPreviews:         queryOption.Preview,
		ReturnShortDescription: true,
	}

	files, err := w.steamCLI.IPublishedFileService().QueryFiles(options)
	if err != nil {
		return publishedfile.FileList{}, err
	}
	return files, nil
}
