package stores

import (
	"backend/models"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type (
	RestaurantStore interface {
		Create(restaurant *models.Restaurant) (string, error)
		Update(restaurant *models.Restaurant) (string, error)
		GetById(id string) (*models.Restaurant, error)
		CreateWithItems(restaurant *models.Restaurant, items *[]models.Item) (string, error)
		DeleteById(id string) error
	}

	restaurantStore struct {
		*gorm.DB
	}
)

func (s *restaurantStore) Create(restaurant *models.Restaurant) (string, error) {
	err := s.DB.Create(&restaurant).Error
	if err != nil {
		log.Error("failed to create restaurant: ", err)
		return "", err
	}

	return restaurant.ID, nil
}

func (s *restaurantStore) Update(restaurant *models.Restaurant) (string, error) {
	err := s.DB.Save(&restaurant).Error
	if err != nil {
		log.Error("failed to create restaurant: ", err)
		return "", err
	}

	return restaurant.ID, nil
}

func (s *restaurantStore) CreateWithItems(restaurant *models.Restaurant, items *[]models.Item) (string, error) {
	restaurant.Items = *items
	err := s.DB.Create(&restaurant).Error
	if err != nil {
		log.Error("failed to create restaurant: ", err)
		return "", err
	}

	return restaurant.ID, nil
}

func (s *restaurantStore) GetById(id string) (*models.Restaurant, error) {
	var restaurant models.Restaurant

	err := s.DB.Preload("Users").Preload("Items").Preload("Categories").Where("id = ? ", id).First(&restaurant).Error

	if err != nil {
		log.Error("can't find restaurant: ", err)
		return nil, err
	}

	return &restaurant, nil
}

func (s *restaurantStore) DeleteById(id string) error {
	return s.DB.Delete(&models.Restaurant{}, "id = ?", id).Error
}
