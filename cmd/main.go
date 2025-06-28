package main

import (
	"log"

	"github.com/sparky1911/auth-service/config"
)

func main() {
	cfg, err := config.LoadConfig("config")
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	log.Printf("Starting %s service in %s mode on port %d\n", cfg.App.Name, cfg.App.Env, cfg.App.Port)

	// TODO: Initialize DB, Redis, routes, etc.
}
