package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type PasswordConfig struct {
	SaltLength int `split_words:"true" json:"PASSWORD_SALT_LENGTH"`
}

var Password *PasswordConfig

func loadPasswordConfig() {
	Password = &PasswordConfig{}
	err := envconfig.Process("password", Password)
	if err != nil {
		log.Fatal(err.Error())
	}
}
