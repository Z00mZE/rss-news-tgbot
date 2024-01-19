package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Telegram struct {
	Token     string `envconfig:"token" required:"true"`
	ChannelID int64  `envconfig:"channel_id" required:"true"`
}

type Publisher struct {
	Interval time.Duration `envconfig:"interval" default:"1m"`
}
type Config struct {
	Telegram  Telegram  `envconfig:"telegram"`
	Publisher Publisher `envconfig:"publisher"`
	LogLevel  string    `envconfig:"log_level" default:"prod"`
}

func NewConfig() (*Config, error) {
	const cfgPrefix = `bot`
	var conf = new(Config)
	if err := envconfig.Process(cfgPrefix, conf); err != nil {
		return nil, err
	}

	return conf, nil
}
