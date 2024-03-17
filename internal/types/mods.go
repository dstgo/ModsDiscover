package types

import "github.com/dstgo/steamapi/types/publishedfile"

type SearchModsOption struct {
	Page int    `query:"page" binding:"gt=0" default:"1"`
	Size int    `query:"size" binding:"gt=0,lte=100" default:"10"`
	Text string `query:"text"`
	Lang int    `query:"lang" binding:"gt=0" default:"6"`
	// return tags
	Tags bool `query:"tags" default:"true"`
	// return previews
	Preview bool `query:"preview" default:"true"`
}

type SearchModsResult struct {
	Total int                  `json:"total"`
	List  []publishedfile.File `json:"list"`
}
