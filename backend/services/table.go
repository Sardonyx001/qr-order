package services

import (
	"backend/models"
	"backend/stores"
)

type (
	TableService interface {
		GetTableById(id string) (*models.Table, error)
		GetTables() ([]models.Table, error)
		CreateTable(table *models.Table) (string, error)
		UpdateTable(table *models.Table, id string) (string, error)
		DeleteTable(id string) error
	}

	tableService struct {
		stores *stores.Stores
	}
)

func NewTableService(stores *stores.Stores) TableService {
	return &tableService{stores}
}

func (s *tableService) GetTableById(id string) (*models.Table, error) {
	return s.stores.Table.GetById(id)
}

func (s *tableService) GetTables() ([]models.Table, error) {
	return s.stores.Table.GetAll()
}

func (s *tableService) CreateTable(table *models.Table) (string, error) {
	tableID, err := s.stores.Table.Create(table)
	return tableID, err
}

func (s *tableService) UpdateTable(newTable *models.Table, id string) (string, error) {
	table, err := s.stores.Table.GetById(id)
	if err != nil {
		return "", err
	}
	table.Name = newTable.Name
	table.Empty = newTable.Empty

	updatedTableID, err := s.stores.Table.Update(table)
	return updatedTableID, err
}

func (s *tableService) DeleteTable(id string) error {
	return s.stores.Table.DeleteById(id)
}
