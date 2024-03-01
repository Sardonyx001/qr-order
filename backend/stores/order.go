package stores

import (
	"backend/models"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type (
	OrderStore interface {
		Create(order *models.Order) (string, error)
		Update(order *models.Order) (string, error)
		GetAll() ([]models.Order, error)
		GetById(id string) (*models.Order, error)
		GetByCustomerId(id string) ([]models.Order, error)
		DeleteById(id string) error
	}

	orderStore struct {
		*gorm.DB
	}
)

func NewOrderStore(db *gorm.DB) OrderStore {
	return &orderStore{db}
}

func (s *orderStore) Create(order *models.Order) (string, error) {
	result := s.DB.Create(order)
	if result.Error != nil {
		return "", result.Error
	}
	return order.ID, nil
}

func (s *orderStore) Update(order *models.Order) (string, error) {
	result := s.DB.Save(order)
	if result.Error != nil {
		return "", result.Error
	}
	return order.ID, nil
}

func (s *orderStore) GetById(id string) (*models.Order, error) {
	var order models.Order
	result := s.DB.Preload("OrderItems").First(&order, "id = ?", id)
	if result.Error != nil {
		log.Error("can't find order: ", result.Error)
		return nil, result.Error
	}
	return &order, nil
}

func (s *orderStore) GetByCustomerId(id string) ([]models.Order, error) {
	var orders []models.Order
	result := s.DB.Preload("OrderItems").Find(&orders).Where("customer_id = ?", id)
	if result.Error != nil {
		log.Error("can't find order: ", result.Error)
		return nil, result.Error
	}
	return orders, nil
}

func (s *orderStore) GetAll() ([]models.Order, error) {
	var orders []models.Order

	result := s.DB.Preload("OrderItems").Find(&orders)
	if result.Error != nil {
		log.Error("Can't find orders: ", result.Error)
		return nil, result.Error
	}
	return orders, nil
}

func (s *orderStore) DeleteById(id string) error {
	return s.DB.Delete(&models.Order{}, "id = ?", id).Error
}
