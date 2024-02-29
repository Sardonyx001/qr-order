package services

import (
	"backend/config"
	"backend/models"
	"backend/stores"

	"golang.org/x/crypto/bcrypt"
)

type (
	AdminService interface {
		GetAdminById(id string) (*models.Admin, error)
		GetAdminByUsername(username string) (*models.Admin, error)
		CreateAdmin(creds *config.BasicAuth) (string, error)
		// UpdateAdminById(user *models.User) (string, error)
		// DeleteAdmin(id string) error
	}

	adminService struct {
		stores *stores.Stores
	}
)

func (s *adminService) CreateAdmin(creds *config.BasicAuth) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(creds.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	admin := models.Admin{
		PasswordHash: string(encryptedPassword),
		Username:     creds.Username,
	}

	adminId, err := s.stores.Admin.Create(&admin)
	return adminId, err
}

func (s *adminService) GetAdminById(id string) (*models.Admin, error) {
	var admin *models.Admin
	admin, err := s.stores.Admin.GetById(id)
	return admin, err
}

func (s *adminService) GetAdminByUsername(username string) (*models.Admin, error) {
	var admin *models.Admin
	admin, err := s.stores.Admin.GetByUsername(username)
	return admin, err
}
