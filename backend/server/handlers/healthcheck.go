package handlers

import (
	s "backend/server"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	HealthCheckHandler interface {
		HealthCheck(c echo.Context) error
	}

	healthCheckHandler struct {
		server *s.Server
	}
)

func (h *healthCheckHandler) HealthCheck(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	db, err := h.server.DB.DB()
	if err != nil {
		h.server.Echo.Logger.Fatal(err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		h.server.Echo.Logger.Fatal("Database down: ", err)
	}

	resp := map[string]string{
		"message": "database ready",
	}

	return c.JSON(http.StatusOK, resp)
}
