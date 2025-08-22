package main

import (
	"github.com/ImHypna/gopportunities.git/config"
	r "github.com/ImHypna/gopportunities.git/router"
)

var (
	logger config.Logger
)

func main() {

	//Initialize Configs
	err := config.Init()
	if err != nil {
		logger.ErrorF("config initialization error: %v", err)
		return
	}

	//Initialize Router

	r.Initialize()
}
