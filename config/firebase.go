package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type FirebaseConfig struct {
	AuthKey string `split_words:"true" json:"FIREBASE_AUTH_KEY"`
}

var Firebase *FirebaseConfig

func loadFirebaseConfig() {
	Firebase = &FirebaseConfig{}
	err := envconfig.Process("firebase", Firebase)
	if err != nil {
		log.Fatal(err.Error())
	}
}
