package main

import (
	"log"
	"myproject/pkg/config"
	"myproject/pkg/di"
)

func main() {
	// Load configuration
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	server, err := di.InitiallizeEvent(conf)
	if err != nil {
		log.Fatal("failed to initialize the files")
	}

	server.Start(conf)
}
