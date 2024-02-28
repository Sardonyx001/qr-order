package handlers

import (
	"backend/config"
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type (
	AuthHandler interface {
		LoginForUser(c echo.Context) error
		LoginForAdmin(c echo.Context) error
		GetUserById(id string) *models.User
	}

	authHandler struct {
		services.AuthService
	}
)

func (h *authHandler) LoginForUser(c echo.Context) error {
	userAuth := new(config.BasicAuth)

	if err := c.Bind(userAuth); err != nil {
		return c.JSON(http.StatusBadRequest, "Request doesn't match schema")
	}

	if err := userAuth.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, "Required fields are empty or not valid")
	}

	accessToken, _, err := h.AuthService.GenerateAccessToken(userAuth.Username, userAuth.Password, false)
	if err != nil {
		log.Error("Failed to authenticate: ", err)
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": accessToken,
	})
}

func (h *authHandler) LoginForAdmin(c echo.Context) error {
	userAuth := new(config.BasicAuth)

	if err := c.Bind(userAuth); err != nil {
		return c.JSON(http.StatusBadRequest, "Request doesn't match schema")
	}

	if err := userAuth.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, "Required fields are empty or not valid")
	}

	accessToken, _, err := h.AuthService.GenerateAccessToken(userAuth.Username, userAuth.Password, true)
	if err != nil {
		log.Error("Failed to authenticate: ", err)
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": accessToken,
	})
}

func (h *authHandler) GetUserFromToken(c echo.Context) *models.User {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*config.JwtCustomClaims)
	return h.AuthService.GetUserById(claims.ID)
}
