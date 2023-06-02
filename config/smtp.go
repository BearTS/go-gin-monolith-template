package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type SmtpConfig struct {
	Host  string `split_words:"true" json:"SMTP_HOST"`
	Port  string `split_words:"true" json:"SMTP_PORT"`
	Email string `split_words:"true" json:"SMTP_EMAIL"`
	Pass  string `split_words:"true" json:"SMTP_PASS"`
}

var Smtp *SmtpConfig

func loadSmtpConfig() {
	Smtp = &SmtpConfig{}
	err := envconfig.Process("smtp", Smtp)
	if err != nil {
		log.Fatal(err.Error())
	}
}
