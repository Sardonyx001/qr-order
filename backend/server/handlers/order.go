package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderHandler interface {
	GetOrders(c echo.Context) error
	GetOrderById(c echo.Context) error
	GetOrdersForCustomer(c echo.Context) error
	CreateOrder(c echo.Context) error
	UpdateOrderById(c echo.Context) error
	DeleteOrderById(c echo.Context) error
}

type orderHandler struct {
	services.OrderService
	services.OrderItemService
}

func (h *orderHandler) GetOrders(c echo.Context) error {
	orders, err := h.OrderService.GetOrders()
	if err != nil {
		return c.JSON(http.StatusNotFound, "No orders found")
	}
	return c.JSON(http.StatusOK, orders)
}

func (h *orderHandler) GetOrderById(c echo.Context) error {
	orderID := c.Param("order_id")
	order, err := h.OrderService.GetOrderById(orderID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "No order found")
	}
	return c.JSON(http.StatusOK, order)
}

func (h *orderHandler) GetOrdersForCustomer(c echo.Context) error {
	customer_id := c.Param("customer_id")

	orders, err := h.OrderService.GetOrderByCustomerId(customer_id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "No order found")
	}
	return c.JSON(http.StatusOK, orders)
}

func (h *orderHandler) CreateOrder(c echo.Context) error {
	order := new(models.Order)
	if err := c.Bind(order); err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to create order")
	}
	orderID, err := h.OrderService.CreateOrder(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create order")
	}

	return c.JSON(http.StatusCreated,
		map[string]string{
			"id": orderID,
		})
}

func (h *orderHandler) UpdateOrderById(c echo.Context) error {
	orderID := c.Param("order_id")
	order := new(models.Order)
	if err := c.Bind(order); err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to update order")
	}
	orderID, err := h.OrderService.UpdateOrder(order, orderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update order")
	}
	return c.JSON(http.StatusOK, map[string]string{"id": orderID})
}

func (h *orderHandler) DeleteOrderById(c echo.Context) error {
	orderID := c.Param("order_id")
	if err := h.OrderService.DeleteOrder(orderID); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete order")
	}
	return c.NoContent(http.StatusNoContent)
}
