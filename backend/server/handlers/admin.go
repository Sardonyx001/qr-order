package handlers

import (
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	AdminHandler interface {
		GetAdminById(c echo.Context) error
		CreateAdmin(c echo.Context) error
		UpdateAdminById(c echo.Context) error
		DeleteAdminById(c echo.Context) error
	}

	adminHandler struct {
		services.AdminService
	}
)

func (h *adminHandler) GetAdminById(c echo.Context) error {
	userID := c.Param("id")
	user, err := h.AdminService.GetAdminById(userID)
	if err != nil {

		return c.JSON(http.StatusNotFound, utils.Error{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *adminHandler) CreateAdmin(c echo.Context) error {
	userAuth := new(utils.BasicAuth)

	if err := c.Bind(userAuth); err != nil {
		return c.JSON(http.StatusBadRequest, "Request doesn't match schema")
	}

	if err := userAuth.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, "Required fields are empty or not valid")
	}

	_, err := h.AdminService.GetAdminByUsername(userAuth.Username)

	if err == nil {
		return c.JSON(http.StatusBadRequest, "User already exists")
	}

	userId, err := h.AdminService.CreateAdmin(userAuth)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Server error")
	}

	return c.JSON(http.StatusCreated, userId)

}

func (h *adminHandler) UpdateAdminById(c echo.Context) error {
	return c.JSON(http.StatusOK, "UpdateUserById")
}

func (h *adminHandler) DeleteAdminById(c echo.Context) error {
	return c.JSON(http.StatusOK, "Called DeleteUserById")
}
