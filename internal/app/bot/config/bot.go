package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Telegram struct {
	Token string `envconfig:"token" required:"true"`
}
type Settings struct {
	Telegram Telegram `envconfig:"telegram"`
}

func NewConfig() (*Settings, error) {
	var conf = new(Settings)
	if err := envconfig.Process("", conf); err != nil {
		return nil, err
	}

	return conf, nil
}
