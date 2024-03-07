package conf

import (
	"github.com/dstgo/tracker/internal/pkg/locale"
	"github.com/dstgo/tracker/pkg/config"
)

// AppConf wilson config contains all needed configurations
type AppConf struct {
	ServerConf *ServerConf  `mapstructure:"app"`
	DataConf   *DataConf    `mapstructure:"data"`
	LogConf    *LogConf     `mapstructure:"log"`
	JwtConf    *JwtConf     `mapstructure:"jwt"`
	LocaleConf *locale.Conf `mapstructure:"locale"`
	EmailConf  *EmailConf   `mapstructure:"email"`
	DstConf    *DstConf     `mapstructure:"dst"`
	BuildMeta  BuildInfo
}

func NewAppConf(config *config.Config, buildInfo BuildInfo) (*AppConf, error) {
	cfg := new(AppConf)
	if err := config.Viper().Unmarshal(cfg); err != nil {
		return nil, err
	}

	if len(buildInfo.Version) == 0 {
		buildInfo.Version = "none"
	}

	if len(buildInfo.Author) == 0 {
		buildInfo.Author = "none"
	}

	if len(buildInfo.BuildTime) == 0 {
		buildInfo.BuildTime = "none"
	}

	cfg.BuildMeta = buildInfo

	return cfg, nil
}
