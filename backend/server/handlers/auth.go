package handlers

import (
	"backend/logger"
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type (
	AuthHandler interface {
		LoginForUser(c echo.Context) error
		LoginForAdmin(c echo.Context) error
	}

	authHandler struct {
		services.AuthService
		services.UserService
		services.AdminService
	}
)

func (h *authHandler) LoginForUser(c echo.Context) error {
	userAuth := new(utils.BasicAuth)

	if err := c.Bind(userAuth); err != nil {
		return c.JSON(http.StatusBadRequest, "Request doesn't match schema")
	}

	if err := userAuth.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, "Required fields are empty or not valid")
	}

	user, err := h.UserService.GetUserByUsername(userAuth.Username)
	if err != nil || (bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userAuth.Password)) != nil) {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	accessToken, _, err := h.AuthService.GenerateAccessToken(user.ID, false)
	if err != nil {
		logger.Error("Failed to authenticate: ", zap.Error(err))
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": accessToken,
	})
}

func (h *authHandler) LoginForAdmin(c echo.Context) error {
	userAuth := new(utils.BasicAuth)

	if err := c.Bind(userAuth); err != nil {
		return c.JSON(http.StatusBadRequest, "Request doesn't match schema")
	}

	if err := userAuth.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, "Required fields are empty or not valid")
	}

	admin, err := h.AdminService.GetAdminByUsername(userAuth.Username)
	if err != nil || (bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(userAuth.Password)) != nil) {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	accessToken, _, err := h.AuthService.GenerateAccessToken(admin.ID, true)
	if err != nil {
		logger.Error("Failed to authenticate: ", zap.Error(err))
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": accessToken,
	})
}
