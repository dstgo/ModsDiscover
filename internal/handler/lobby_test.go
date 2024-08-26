package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/dstgo/tracker/internal/assets"
	"github.com/dstgo/tracker/internal/data"
	"github.com/dstgo/tracker/pkg/lobbyapi"
	"github.com/go-resty/resty/v2"
	"testing"
)

func TestLobbyMongoHandler_GetAllLobbyServers(t *testing.T) {
	geoip, err := data.LoadGeoIpDBInMem(assets.GeopIp2CityDB)
	assert.Nil(t, err)

	client := lobbyapi.New("")

	handler := LobbyMongoHandler{geoip: geoip, lobby: client}

	servers, err := handler.GetAllServersFromLobby(context.Background(), 30, 0)
	assert.Nil(t, err)

	t.Log(len(servers))
}

func TestLobbyMongoHandler_GetAllLobbyServersWithProxy(t *testing.T) {
	geoip, err := data.LoadGeoIpDBInMem(assets.GeopIp2CityDB)
	assert.Nil(t, err)

	c := resty.New()
	c.SetProxy("http://127.0.0.1:7890")
	client := lobbyapi.NewWith("", c)

	handler := LobbyMongoHandler{geoip: geoip, lobby: client}

	servers, err := handler.GetAllServersFromLobby(context.Background(), 30, 0)
	assert.Nil(t, err)

	t.Log(len(servers))
}
