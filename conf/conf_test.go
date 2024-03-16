package conf

import (
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/henrylee2cn/goutil/calendar/cron"
	"testing"
	"time"
)

func TestLoadConf(t *testing.T) {
	conf, err := Load("conf.yaml")
	assert.Nil(t, err)
	t.Log(conf)
}

func TestCronParse(t *testing.T) {
	parse, err := cron.ParseStandard("0 3 * * *")
	assert.Nil(t, err)
	t.Log(parse.Next(time.Now()))
}
