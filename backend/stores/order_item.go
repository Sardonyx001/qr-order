package stores

import (
	"backend/models"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type (
	OrderItemStore interface {
		Create(orderItem *models.OrderItem) (string, error)
		Update(orderItem *models.OrderItem) (string, error)
		GetAll() ([]models.OrderItem, error)
		GetById(id string) (*models.OrderItem, error)
		DeleteById(id string) error
	}

	orderItemStore struct {
		*gorm.DB
	}
)

func NewOrderItemStore(db *gorm.DB) OrderItemStore {
	return &orderItemStore{db}
}

func (s *orderItemStore) Create(orderItem *models.OrderItem) (string, error) {
	result := s.DB.Create(orderItem)
	if result.Error != nil {
		return "", result.Error
	}
	return orderItem.ID, nil
}

func (s *orderItemStore) Update(orderItem *models.OrderItem) (string, error) {
	result := s.DB.Save(orderItem)
	if result.Error != nil {
		return "", result.Error
	}
	return orderItem.ID, nil
}

func (s *orderItemStore) GetById(id string) (*models.OrderItem, error) {
	var orderItem models.OrderItem
	result := s.DB.First(&orderItem, "id = ?", id)
	if result.Error != nil {
		log.Error("can't find order item: ", result.Error)
		return nil, result.Error
	}
	return &orderItem, nil
}

func (s *orderItemStore) GetAll() ([]models.OrderItem, error) {
	var orderItems []models.OrderItem

	result := s.DB.Find(&orderItems)
	if result.Error != nil {
		log.Error("Can't find order items: ", result.Error)
		return nil, result.Error
	}
	return orderItems, nil
}

func (s *orderItemStore) DeleteById(id string) error {
	return s.DB.Delete(&models.OrderItem{}, "id = ?", id).Error
}
