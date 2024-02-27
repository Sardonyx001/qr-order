package stores

import (
	"backend/models"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type (
	UserStore interface {
		Create(user *models.User) (string, error)
		GetById(id string) (*models.User, error)
		DeleteById(id string) error
	}

	userStore struct {
		*gorm.DB
	}
)

func (s *userStore) Create(user *models.User) (string, error) {
	err := s.DB.Create(&user).Error

	if err != nil {
		log.Error("failed to create user", err)
		return "", err
	}

	return user.ID, nil
}

func (s *userStore) GetById(id string) (*models.User, error) {
	var user models.User

	err := s.DB.Where("id = ? ", id).Take(&user).Error

	if err != nil {
		log.Error("can't find user", err)
		return nil, err
	}

	return &user, nil
}

func (s *userStore) DeleteById(id string) error {
	return s.DB.Delete(&models.User{}, id).Error
}
