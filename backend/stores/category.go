package stores

import (
	"backend/models"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type (
	CategoryStore interface {
		Create(category *models.Category) error
		GetAll() ([]models.Category, error)
		GetById(id string) (*models.Category, error)
		GetByName(name string) (*models.Category, error)
		Update(category *models.Category) error
		DeleteById(id string) error
	}

	categoryStore struct {
		*gorm.DB
	}
)

func (s *categoryStore) Create(category *models.Category) error {
	return s.DB.Create(category).Error
}

func (s *categoryStore) GetAll() ([]models.Category, error) {
	var categories []models.Category

	result := s.DB.Preload("Items").Find(&categories)
	if result.Error != nil {
		log.Error("can't find category ", result.Error)
		return nil, result.Error
	}
	return categories, nil
}

func (s *categoryStore) GetById(id string) (*models.Category, error) {
	var category models.Category
	result := s.DB.Preload("Items").First(&category, "id = ?", id)
	if result.Error != nil {
		log.Error("can't find category ", result.Error)
		return nil, result.Error
	}
	return &category, nil
}

func (s *categoryStore) GetByName(name string) (*models.Category, error) {
	var category models.Category
	result := s.DB.Preload("Restaurants").Preload("Items").First(&category, "name = ?", name)
	if result.Error != nil {
		log.Error("can't find category ", result.Error)
		return nil, result.Error
	}
	return &category, nil
}

func (s *categoryStore) Update(category *models.Category) error {
	return s.DB.Save(category).Error
}

func (s *categoryStore) DeleteById(id string) error {
	return s.DB.Delete(&models.Category{}, "id = ?", id).Error
}
