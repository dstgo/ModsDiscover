package assets

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/oschwald/geoip2-golang"
	"net"
	"testing"
)

func TestGeoip2(t *testing.T) {
	bytes, err := FS.ReadFile("geoip2/GeoLite2-City.mmdb")
	assert.Nil(t, err)

	db, err := geoip2.FromBytes(bytes)
	assert.Nil(t, err)

	city, err := db.City(net.ParseIP("125.227.86.48"))
	assert.Nil(t, err)
	t.Log(city.Country.IsoCode)
	t.Log(city.City.Names)
	t.Log(city.Location.TimeZone)

	fmt.Printf("%s-%s-%s", city.Continent.Names["en"], city.Country.Names["en"], city.City.Names["en"])
}

func BenchmarkGeoip2(b *testing.B) {
	b.ReportAllocs()

	bytes, err := FS.ReadFile("geoip2/GeoLite2-City.mmdb")
	assert.Nil(b, err)

	db, err := geoip2.FromBytes(bytes)
	assert.Nil(b, err)

	for i := 0; i < b.N; i++ {
		db.Country(net.ParseIP("125.227.86.48"))
	}
}
