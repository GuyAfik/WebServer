package main

import (
	"log"
	"github.com/WebServer/server/internal/app"
	"github.com/WebServer/server/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.RunServer(cfg)
}