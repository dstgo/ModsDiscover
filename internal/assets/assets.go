package assets

import "embed"

//go:embed *
var FS embed.FS

const (
	GeopIp2CityDB = "geoip2/GeoLite2-City.mmdb"
)
