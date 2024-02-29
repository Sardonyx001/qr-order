package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	ItemHandler interface {
		GetItemById(c echo.Context) error
		GetItems(c echo.Context) error
		CreateItem(c echo.Context) error
		UpdateItemById(c echo.Context) error
		DeleteItemById(c echo.Context) error
	}

	itemHandler struct {
		services.ItemService
	}
)

func (h *itemHandler) GetItems(c echo.Context) error {
	items, err := h.ItemService.GetItems()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, items)
}

func (h *itemHandler) GetItemById(c echo.Context) error {
	itemID := c.Param("item_id")

	item, err := h.ItemService.GetItemById(itemID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, item)
}

func (h *itemHandler) CreateItem(c echo.Context) error {
	item := new(models.Item)
	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	itemID, err := h.ItemService.CreateItem(item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"id": itemID})
}

func (h *itemHandler) UpdateItemById(c echo.Context) error {
	itemID := c.Param("item_id")

	item := new(models.Item)
	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	itemID, err := h.ItemService.UpdateItem(item, itemID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"id": itemID})
}

func (h *itemHandler) DeleteItemById(c echo.Context) error {
	itemID := c.Param("item_id")

	if err := h.ItemService.DeleteItem(itemID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
