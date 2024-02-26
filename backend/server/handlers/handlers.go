package handlers

import (
	s "backend/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	HelloHandler interface {
		Hello(c echo.Context) error
	}

	helloHandler struct {
		server *s.Server
	}
)

func (h *helloHandler) Hello(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

type Handlers struct {
	HelloHandler
	HealthCheckHandler
}

func New(server *s.Server) *Handlers {
	return &Handlers{
		HelloHandler:       &helloHandler{server: server},
		HealthCheckHandler: &healthCheckHandler{server: server},
	}
}
