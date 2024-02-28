package services

import (
	"backend/models"
	"backend/stores"
)

type (
	RestaurantService interface {
		// GetRestaurantById(id string) (*models.Restaurant, error)
		CreateRestaurant(r *models.Restaurant) (string, error)
		// CreateRestaurantWithItems(r *models.Restaurant, is *[]models.Item) (string, error)
		DeleteRestaurant(id string) error
	}

	restaurantService struct {
		stores *stores.Stores
	}
)

func (s *restaurantService) CreateRestaurant(restaurant *models.Restaurant) (string, error) {
	restaurant_id, err := s.stores.Restaurant.Create(restaurant)
	return restaurant_id, err
}

func (s *restaurantService) DeleteRestaurant(id string) error {
	err := s.stores.Restaurant.DeleteById(id)
	return err
}
