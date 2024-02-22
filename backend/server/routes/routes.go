package routes

import (
	s "backend/server"
	"backend/server/handlers"

	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(server *s.Server) {
	helloHandler := handlers.NewHelloHandlers(server)

	server.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
		Output: server.Echo.Logger.Output(),
	}))

	server.Echo.GET("/hello", helloHandler.HelloHandler)
	server.Echo.GET("/healthcheck", helloHandler.HealthcheckHandler)
}
