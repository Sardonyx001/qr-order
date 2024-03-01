package services

import (
	"backend/models"
	"backend/stores"
)

type CustomerService interface {
	GetCustomerById(id string) (*models.Customer, error)
	GetCustomers() ([]models.Customer, error)
	CreateCustomer(customer *models.Customer) (string, error)
	UpdateCustomer(customer *models.Customer, id string) (string, error)
	DeleteCustomer(id string) error
}

type customerService struct {
	stores *stores.Stores
}

func (s *customerService) GetCustomerById(id string) (*models.Customer, error) {
	return s.stores.Customer.GetById(id)
}

func (s *customerService) GetCustomers() ([]models.Customer, error) {
	return s.stores.Customer.GetAll()
}

func (s *customerService) CreateCustomer(customer *models.Customer) (string, error) {
	return s.stores.Customer.Create(customer)
}

func (s *customerService) UpdateCustomer(customer *models.Customer, id string) (string, error) {
	return s.stores.Customer.Update(customer)
}

func (s *customerService) DeleteCustomer(id string) error {
	return s.stores.Customer.DeleteById(id)
}
