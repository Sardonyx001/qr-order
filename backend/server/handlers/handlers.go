package handlers

import (
	s "backend/server"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type HelloHandlers struct {
	server *s.Server
}

func NewHelloHandlers(server *s.Server) *HelloHandlers {
	return &HelloHandlers{server: server}
}

func (p *HelloHandlers) HelloHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (p *HelloHandlers) HealthcheckHandler(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	db, err := p.server.DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Database down: ", err)
	}

	resp := map[string]string{
		"message": "database ready",
	}

	return c.JSON(http.StatusOK, resp)
}
