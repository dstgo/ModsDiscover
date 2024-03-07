package dst

import "github.com/dstgo/steamapi/types/publishedfile"

type SearchModsOption struct {
	page int
	size int
	text string
	lang string
	tags string
}

type SearchModsResult struct {
	Total int                  `json:"total"`
	List  []publishedfile.File `json:"list"`
}

type QueryModOption struct {
	Id      string `uri:"id" form:"id"`
	Version string `uri:"version" form:"version"`
}

type ModsDetail = publishedfile.File
