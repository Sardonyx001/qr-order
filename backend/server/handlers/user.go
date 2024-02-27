package handlers

import (
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	UserHanlder interface {
		GetUserById(c echo.Context) error
		CreateUser(c echo.Context) error
		UpdateUserById(c echo.Context) error
		DeleteUserById(c echo.Context) error
	}

	userHandler struct {
		services.UserService
	}
)

func (h *userHandler) GetUserById(c echo.Context) error {
	userID := c.Param("id")
	user, err := h.UserService.GetUserById(userID)
	if err != nil {
        
		return c.JSON(http.StatusNotFound, utils.Error{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) CreateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "CreateUser")
}

func (h *userHandler) UpdateUserById(c echo.Context) error {
	return c.JSON(http.StatusOK, "UpdateUserById")
}

func (h *userHandler) DeleteUserById(c echo.Context) error {
	return c.JSON(http.StatusOK, "Called DeleteUserById")
}
