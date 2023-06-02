package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type TokenConfig struct {
	AccessSecret  string `split_words:"true" json:"TOKEN_ACCESS_SECRET"`
	RefreshSecret string `split_words:"true" json:"TOKEN_REFRESH_SECRET"`
}

var Token *TokenConfig

func loadTokenConfig() {
	Token = &TokenConfig{}
	err := envconfig.Process("token", Token)
	if err != nil {
		log.Fatal(err.Error())
	}
}
