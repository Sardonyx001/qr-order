package services

import (
	"backend/models"
	"backend/stores"
)

type (
	OrderItemService interface {
		GetOrderItemById(id string) (*models.OrderItem, error)
		GetOrderItems() ([]models.OrderItem, error)
		CreateOrderItem(orderItem *models.OrderItem) (string, error)
		UpdateOrderItem(orderItem *models.OrderItem, id string) (string, error)
		DeleteOrderItem(id string) error
	}

	orderItemService struct {
		stores *stores.Stores
	}
)

func (s *orderItemService) GetOrderItemById(id string) (*models.OrderItem, error) {
	return s.stores.OrderItem.GetById(id)
}

func (s *orderItemService) GetOrderItems() ([]models.OrderItem, error) {
	return s.stores.OrderItem.GetAll()
}

func (s *orderItemService) CreateOrderItem(orderItem *models.OrderItem) (string, error) {
	orderItemID, err := s.stores.OrderItem.Create(orderItem)
	return orderItemID, err
}

func (s *orderItemService) UpdateOrderItem(orderItem *models.OrderItem, id string) (string, error) {
	oldOrderItem, err := s.stores.OrderItem.GetById(id)
	if err != nil {
		return "", err
	}

	oldOrderItem.Status = orderItem.Status

	updatedOrderItemID, err := s.stores.OrderItem.Update(oldOrderItem)
	return updatedOrderItemID, err
}

func (s *orderItemService) DeleteOrderItem(id string) error {
	err := s.stores.OrderItem.DeleteById(id)
	return err
}
