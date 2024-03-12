package conf

import (
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"testing"
)

func TestLoadConf(t *testing.T) {
	conf, err := Load("conf.yaml")
	assert.Nil(t, err)
	t.Log(conf)
}
