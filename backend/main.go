package main

import (
	config "backend/config"
	"backend/server"
	"log"
)

func main() {
	cfg := config.NewConfig()

	app := server.NewServer(cfg)

	server.ConfigureRoutes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
