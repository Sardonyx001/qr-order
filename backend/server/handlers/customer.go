package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomerHandler interface {
	GetCustomers(c echo.Context) error
	GetCustomerById(c echo.Context) error
	CreateCustomer(c echo.Context) error
	DeleteCustomerById(c echo.Context) error
}

type customerHandler struct {
	services.CustomerService
}

func (h *customerHandler) GetCustomers(c echo.Context) error {
	customers, err := h.CustomerService.GetCustomers()
	if err != nil {
		return c.JSON(http.StatusNotFound, "No customers found")
	}
	return c.JSON(http.StatusOK, customers)
}

func (h *customerHandler) GetCustomerById(c echo.Context) error {
	customerID := c.Param("customer_id")
	customer, err := h.CustomerService.GetCustomerById(customerID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "No customer found")
	}
	return c.JSON(http.StatusOK, customer)
}

func (h *customerHandler) CreateCustomer(c echo.Context) error {
	customer := new(models.Customer)
	if err := c.Bind(customer); err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to create customer")
	}
	id, err := h.CustomerService.CreateCustomer(customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create customer")
	}
	return c.JSON(http.StatusCreated, map[string]string{
		"id": id,
	})
}

func (h *customerHandler) DeleteCustomerById(c echo.Context) error {
	customerID := c.Param("customer_id")
	if err := h.CustomerService.DeleteCustomer(customerID); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete customer")
	}
	return c.JSON(http.StatusOK, "Record deletion succeeded")
}
