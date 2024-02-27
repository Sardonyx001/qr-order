package services

import (
	"backend/models"
	"backend/stores"
)

type (
	UserService interface {
		GetUserById(id string) (*models.User, error)
		// CreateUser(user *models.User) (string, error)
		// UpdateUserById(user *models.User) (string, error)
		// DeleteUser(id string) error
	}

	userService struct {
		stores *stores.Stores
	}
)

func (s *userService) GetUserById(id string) (*models.User, error) {
	var user *models.User
	user, err := s.stores.User.GetById(id)
	return user, err
}
