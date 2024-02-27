package stores

import (
	"backend/models"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type (
	ItemStore interface {
		Create(item *models.Item) (string, error)
		GetById(id string) (*models.Item, error)
		DeleteById(id string) error
	}

	itemStore struct {
		*gorm.DB
	}
)

func (s *itemStore) Create(item *models.Item) (string, error) {
	result := s.DB.Create(&item)
	if result.Error != nil {
		return "", result.Error
	}
	return item.ID, nil
}

func (s *itemStore) GetById(id string) (*models.Item, error) {
	var item models.Item
	result := s.DB.First(&item, id)
	if result.Error != nil {
		log.Error("can't find item: ", result.Error)
		return nil, result.Error
	}
	return &item, nil
}

func (s *itemStore) DeleteById(id string) error {
	return s.DB.Delete(&models.Item{}, id).Error
}