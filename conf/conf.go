package conf

import (
	"github.com/spf13/viper"
	"time"
)

type AppConf struct {
	Http HttpConf `mapstructure:"http"`
	Log  LogConf  `mapstructure:"log"`
	DB   DBConf   `mapstructure:"db"`
	Dst  DstConf  `mapstructure:"dst"`
}

type HttpConf struct {
	Listen       string        `mapstructure:"listen"`
	BasePath     string        `mapstructure:"base"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
	IdleTimeout  time.Duration `mapstructure:"idleTimeout"`
	CertFile     string        `mapstructure:"cert"`
	KeyFile      string        `mapstructure:"key"`
	CacheTTL     time.Duration `mapstructure:"cacheTTL"`
}

type LogConf struct {
	File  string `mapstructure:"file"`
	Level string `mapstructure:"level"`
}

type DBConf struct {
	Address  string `mapstructure:"address"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DataBase string `mapstructure:"database"`
	Params   string `mapstructure:"params"`
}

type DstConf struct {
	SteamKey  string `mapstructure:"steamKey"`
	KleiToken string `mapstructure:"kleiToken"`
	ProxyURL  string `mapstructure:"proxyURL"`

	Lobby LobbyConf `mapstructure:"lobby"`
}

type LobbyConf struct {
	CollectCron string        `mapstructure:"collect"`
	ClearCron   string        `mapstructure:"clear"`
	TTL         time.Duration `mapstructure:"ttl"`
	Timeout     time.Duration `mapstructure:"timeout"`
}

// Load tries to load config file and unmarshal it to *AppConf
func Load(file string) (*AppConf, error) {
	v := viper.New()
	v.SetConfigFile(file)
	var conf AppConf
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := v.Unmarshal(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
