package stores

import (
	"backend/models"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type (
	AdminStore interface {
		Create(admin *models.Admin) (string, error)
		GetById(id string) (*models.Admin, error)
		GetDefaultAdmin(defaultAdmin *models.Admin) error
		GetByUsername(username string) (*models.Admin, error)
		DeleteById(id string) error
	}

	adminStore struct {
		*gorm.DB
	}
)

func (s *adminStore) Create(admin *models.Admin) (string, error) {
	err := s.DB.Create(&admin).Error
	if err != nil {
		log.Error("failed to create admin: ", err)
		return "", err
	}

	return admin.ID, nil
}

func (s *adminStore) GetById(id string) (*models.Admin, error) {
	var admin models.Admin

	err := s.DB.Where("id = ? ", id).Take(&admin).Error

	if err != nil {
		log.Error("can't find admin: ", err)
		return nil, err
	}

	return &admin, nil
}

func (s *adminStore) GetByUsername(username string) (*models.Admin, error) {
	var admin models.Admin

	err := s.DB.Where("username = ? ", username).First(&admin).Error

	if err != nil {
		log.Error("can't find admin ", err)
		return nil, err
	}

	return &admin, nil
}

func (s *adminStore) GetDefaultAdmin(defaultAdmin *models.Admin) error {
	return s.DB.First(&defaultAdmin).Error
}

func (s *adminStore) DeleteById(id string) error {
	return s.DB.Delete(&models.Admin{}, id).Error
}
