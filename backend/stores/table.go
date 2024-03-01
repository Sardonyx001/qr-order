package stores

import (
	"backend/models"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type (
	TableStore interface {
		Create(table *models.Table) (string, error)
		Update(table *models.Table) (string, error)
		GetAll() ([]models.Table, error)
		GetById(id string) (*models.Table, error)
		DeleteById(id string) error
	}

	tableStore struct {
		*gorm.DB
	}
)

func (s *tableStore) Create(table *models.Table) (string, error) {
	result := s.DB.Create(table)
	if result.Error != nil {
		return "", result.Error
	}
	return table.ID, nil
}

func (s *tableStore) Update(table *models.Table) (string, error) {

	result := s.DB.Save(table)
	if result.Error != nil {
		return "", result.Error
	}
	return table.ID, nil
}

func (s *tableStore) GetById(id string) (*models.Table, error) {
	var table models.Table
	result := s.DB.Preload("Orders").First(&table, "id = ?", id)
	if result.Error != nil {
		log.Error("can't find table: ", result.Error)
		return nil, result.Error
	}
	return &table, nil
}

func (s *tableStore) GetAll() ([]models.Table, error) {
	var tables []models.Table

	result := s.DB.Preload("Orders").Find(&tables)
	if result.Error != nil {
		log.Error("Can't find tables: ", result.Error)
		return nil, result.Error
	}
	return tables, nil
}

func (s *tableStore) DeleteById(id string) error {
	return s.DB.Delete(&models.Table{}, id).Error
}
