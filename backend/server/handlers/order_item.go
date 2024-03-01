package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderItemHandler interface {
	GetOrderItems(c echo.Context) error
	GetOrderItemById(c echo.Context) error
	CreateOrderItem(c echo.Context) error
	UpdateOrderItemById(c echo.Context) error
	DeleteOrderItemById(c echo.Context) error
}

type orderItemHandler struct {
	services.OrderItemService
}

func (h *orderItemHandler) GetOrderItems(c echo.Context) error {
	orderItems, err := h.OrderItemService.GetOrderItems()
	if err != nil {
		return c.JSON(http.StatusNotFound, "No order items found")
	}
	return c.JSON(http.StatusOK, orderItems)
}

func (h *orderItemHandler) GetOrderItemById(c echo.Context) error {
	orderItemID := c.Param("order_item_id")
	orderItem, err := h.OrderItemService.GetOrderItemById(orderItemID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "No order item found")
	}
	return c.JSON(http.StatusOK, orderItem)
}

func (h *orderItemHandler) CreateOrderItem(c echo.Context) error {
	orderItem := new(models.OrderItem)
	if err := c.Bind(orderItem); err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to create order item")
	}
	orderItemID, err := h.OrderItemService.CreateOrderItem(orderItem)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create order item")
	}
	return c.JSON(http.StatusCreated, map[string]string{"id": orderItemID})
}

func (h *orderItemHandler) UpdateOrderItemById(c echo.Context) error {
	orderItemID := c.Param("order_item_id")
	orderItem := new(models.OrderItem)
	if err := c.Bind(orderItem); err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to update order item")
	}

	orderItemID, err := h.OrderItemService.UpdateOrderItem(orderItem, orderItemID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update order item")
	}
	return c.JSON(http.StatusOK, map[string]string{"id": orderItemID})
}

func (h *orderItemHandler) DeleteOrderItemById(c echo.Context) error {
	orderItemID := c.Param("order_item_id")
	if err := h.OrderItemService.DeleteOrderItem(orderItemID); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete order item")
	}
	return c.NoContent(http.StatusNoContent)
}
