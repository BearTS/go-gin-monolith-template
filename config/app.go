package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Port     int
	Env      string
	RootPath string
	Migrate  int
}

var App *AppConfig

func loadAppConfig() {
	App = &AppConfig{}
	err := envconfig.Process("app", App)

	_, b, _, _ := runtime.Caller(0)
	App.RootPath = filepath.Join(filepath.Dir(b), "../")

	if err != nil {
		log.Fatal(err.Error())
	}
}
