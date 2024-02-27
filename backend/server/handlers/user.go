package handlers

import (
	"backend/auth"
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	UserHandler interface {
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
	userAuth := new(auth.BasicAuth)

	if err := c.Bind(userAuth); err != nil {
		return c.JSON(http.StatusBadRequest, "Request doesn't match schema")
	}

	if err := userAuth.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, "Required fields are empty or not valid")
	}

	_, err := h.UserService.GetUserByUsername(userAuth.Username)

	if err == nil {
		return c.JSON(http.StatusBadRequest, "User already exists")
	}

	userId, err := h.UserService.CreateUser(userAuth)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Server error occured")
	}

	return c.JSON(http.StatusCreated, userId)

}

func (h *userHandler) UpdateUserById(c echo.Context) error {
	return c.JSON(http.StatusOK, "UpdateUserById")
}

func (h *userHandler) DeleteUserById(c echo.Context) error {
	return c.JSON(http.StatusOK, "Called DeleteUserById")
}
