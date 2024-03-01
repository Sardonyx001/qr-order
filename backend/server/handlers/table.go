package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	TableHandler interface {
		GetTables(c echo.Context) error
		GetTableById(c echo.Context) error
		CreateTable(c echo.Context) error
		UpdateTableById(c echo.Context) error
		DeleteTableById(c echo.Context) error
	}

	tableHandler struct {
		tableService services.TableService
	}
)

func NewTableHandler(tableService services.TableService) TableHandler {
	return &tableHandler{tableService}
}

func (h *tableHandler) GetTables(c echo.Context) error {
	tables, err := h.tableService.GetTables()
	if err != nil {
		return c.JSON(http.StatusNotFound, "No tables found")
	}

	return c.JSON(http.StatusOK, tables)
}

func (h *tableHandler) GetTableById(c echo.Context) error {
	tableID := c.Param("table_id")

	table, err := h.tableService.GetTableById(tableID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "No table found")
	}

	return c.JSON(http.StatusOK, table)
}

func (h *tableHandler) CreateTable(c echo.Context) error {
	table := models.Table{}
	if err := c.Bind(&table); err != nil {
		return c.JSON(http.StatusBadRequest, "Record creation failed")
	}

	tableID, err := h.tableService.CreateTable(&table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Record creation failed")
	}

	return c.JSON(http.StatusCreated, map[string]string{"id": tableID})
}

func (h *tableHandler) UpdateTableById(c echo.Context) error {
	tableID := c.Param("table_id")

	table := new(models.Table)
	if err := c.Bind(table); err != nil {
		return c.JSON(http.StatusBadRequest, "Record creation failed")
	}

	tableID, err := h.tableService.UpdateTable(table, tableID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Record creation failed")
	}

	return c.JSON(http.StatusOK, map[string]string{"id": tableID})
}

func (h *tableHandler) DeleteTableById(c echo.Context) error {
	tableID := c.Param("table_id")

	if err := h.tableService.DeleteTable(tableID); err != nil {
		return c.JSON(http.StatusInternalServerError, "Record deletion failed")
	}
	
	return c.JSON(http.StatusOK, "Record deletion succeeded")
}
