package services

import (
	"backend/models"
	"backend/stores"
)

type (
	ItemService interface {
		GetItemById(id string) (*models.Item, error)
		GetItems() ([]models.Item, error)
		CreateItem(r *models.Item) (string, error)
		UpdateItem(r *models.Item, id string) (string, error)
		DeleteItem(id string) error
	}

	itemService struct {
		stores *stores.Stores
	}
)

func (s *itemService) GetItemById(id string) (*models.Item, error) {
	return s.stores.Item.GetById(id)
}

func (s *itemService) GetItems() ([]models.Item, error) {
	return s.stores.Item.GetAll()
}

func (s *itemService) CreateItem(item *models.Item) (string, error) {
	item_id, err := s.stores.Item.Create(item)
	return item_id, err
}
func (s *itemService) UpdateItem(item *models.Item, id string) (string, error) {
	oldItem, err := s.stores.Item.GetById(id)
	if err != nil {
		return "", err
	}
	oldItem.Name = item.Name
	oldItem.Options = item.Options
	oldItem.Name = item.Name

	item_id, err := s.stores.Item.Update(oldItem)
	return item_id, err
}

func (s *itemService) DeleteItem(id string) error {
	err := s.stores.Item.DeleteById(id)
	return err
}
