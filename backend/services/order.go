package services

import (
	"backend/models"
	"backend/stores"
)

type (
	OrderService interface {
		GetOrderById(id string) (*models.Order, error)
		GetOrderByCustomerId(id string) ([]models.Order, error)
		GetOrders() ([]models.Order, error)
		CreateOrder(order *models.Order) (string, error)
		UpdateOrder(order *models.Order, id string) (string, error)
		DeleteOrder(id string) error
	}

	orderService struct {
		stores *stores.Stores
	}
)

func (s *orderService) GetOrderById(id string) (*models.Order, error) {
	return s.stores.Order.GetById(id)
}

func (s *orderService) GetOrderByCustomerId(id string) ([]models.Order, error) {
	return s.stores.Order.GetByCustomerId(id)
}

func (s *orderService) GetOrders() ([]models.Order, error) {
	return s.stores.Order.GetAll()
}

func (s *orderService) CreateOrder(order *models.Order) (string, error) {
	orderID, err := s.stores.Order.Create(order)
	return orderID, err
}

func (s *orderService) UpdateOrder(order *models.Order, id string) (string, error) {
	updatedOrderID, err := s.stores.Order.Update(order)
	return updatedOrderID, err
}

func (s *orderService) DeleteOrder(id string) error {
	err := s.stores.Order.DeleteById(id)
	return err
}
