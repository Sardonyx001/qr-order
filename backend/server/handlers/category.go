package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	CategoryHandler interface {
		GetCategoryById(c echo.Context) error
		CreateCategory(c echo.Context) error
		UpdateCategoryById(c echo.Context) error
		DeleteCategoryById(c echo.Context) error
	}

	categoryHandler struct {
		services.CategoryService
	}
)

func (h *categoryHandler) GetCategoryById(c echo.Context) error {
	categoryID := c.Param("category_id")

	category, err := h.CategoryService.GetCategoryById(categoryID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, category)
}

func (h *categoryHandler) CreateCategory(c echo.Context) error {
	category := models.Category{}
	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	categoryID, err := h.CategoryService.CreateCategory(&category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"id": categoryID})
}

func (h *categoryHandler) UpdateCategoryById(c echo.Context) error {
	categoryID := c.Param("category_id")

	category := new(models.Category)
	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	categoryID, err := h.CategoryService.UpdateCategory(category, categoryID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"id": categoryID})
}

func (h *categoryHandler) DeleteCategoryById(c echo.Context) error {
	categoryID := c.Param("category_id")

	if err := h.CategoryService.DeleteCategory(categoryID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
