package main

import (
	config "backend/config"
	"backend/server"
	"backend/server/routes"
	"log"
)

func main() {
	cfg := config.NewConfig()

	app := server.NewServer(cfg)

	routes.ConfigureRoutes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
