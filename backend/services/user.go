package services

import (
	"backend/config"
	"backend/models"
	"backend/stores"

	"golang.org/x/crypto/bcrypt"
)

type (
	UserService interface {
		GetUserById(id string) (*models.User, error)
		GetUserByUsername(username string) (*models.User, error)
		CreateUser(creds *config.BasicAuth) (string, error)
		DeleteUser(id string) error
	}

	userService struct {
		stores *stores.Stores
	}
)

func (s *userService) CreateUser(creds *config.BasicAuth) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(creds.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}
	defaultAdmin := models.Admin{}
	result := s.stores.DB.First(&defaultAdmin)

	if result.Error != nil {
		return "", result.Error
	}

	user := models.User{
		Username:     creds.Username,
		PasswordHash: string(encryptedPassword),
		Admin:        defaultAdmin,
	}
	user.PasswordHash = string(encryptedPassword)
	user.Username = creds.Username

	userId, err := s.stores.User.Create(&user)
	return userId, err
}

func (s *userService) GetUserById(id string) (*models.User, error) {
	var user *models.User
	user, err := s.stores.User.GetById(id)
	return user, err
}

func (s *userService) GetUserByUsername(username string) (*models.User, error) {
	var user *models.User
	user, err := s.stores.User.GetByUsername(username)
	return user, err
}

func (s *userService) DeleteUser(id string) error {
	return s.stores.User.DeleteById(id)
}
