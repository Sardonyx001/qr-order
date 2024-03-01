package stores

import (
	"backend/models"

	"gorm.io/gorm"
)

type CustomerStore interface {
	Create(customer *models.Customer) (string, error)
	Update(customer *models.Customer) (string, error)
	GetAll() ([]models.Customer, error)
	GetById(id string) (*models.Customer, error)
	DeleteById(id string) error
}

type customerStore struct {
	db *gorm.DB
}

func NewCustomerStore(db *gorm.DB) CustomerStore {
	return &customerStore{db}
}

func (s *customerStore) Create(customer *models.Customer) (string, error) {
	result := s.db.Create(&customer)
	if result.Error != nil {
		return "", result.Error
	}
	return customer.ID, nil
}

func (s *customerStore) Update(customer *models.Customer) (string, error) {
	result := s.db.Save(&customer)
	if result.Error != nil {
		return "", result.Error
	}
	return customer.ID, nil
}

func (s *customerStore) GetById(id string) (*models.Customer, error) {
	var customer models.Customer
	result := s.db.First(&customer, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func (s *customerStore) GetAll() ([]models.Customer, error) {
	var customers []models.Customer
	result := s.db.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func (s *customerStore) DeleteById(id string) error {
	result := s.db.Delete(&models.Customer{}, id)
	return result.Error
}
