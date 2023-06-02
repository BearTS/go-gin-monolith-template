package main

import (
	"github.com/BearTS/go-gin-monolith/app"
	"github.com/BearTS/go-gin-monolith/config"
)

func main() {
	config.LoadConfigs()
	app.App()
}
