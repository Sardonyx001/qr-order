package services

import (
	"backend/models"
	"backend/stores"
)

type (
	CategoryService interface {
		GetCategories() ([]models.Category, error)
		GetCategoryById(id string) (*models.Category, error)
		GetCategoryByName(name string) (*models.Category, error)
		CreateCategory(category *models.Category) (string, error)
		UpdateCategory(newCategory *models.Category, id string) (string, error)
		DeleteCategory(id string) error
	}

	categoryService struct {
		stores *stores.Stores
	}
)

func (s *categoryService) GetCategories() ([]models.Category, error) {
	return s.stores.Category.GetAll()
}

func (s *categoryService) GetCategoryById(id string) (*models.Category, error) {
	return s.stores.Category.GetById(id)
}

func (s *categoryService) GetCategoryByName(name string) (*models.Category, error) {
	return s.stores.Category.GetByName(name)
}

func (s *categoryService) CreateCategory(category *models.Category) (string, error) {
	err := s.stores.Category.Create(category)
	return category.ID, err
}

func (s *categoryService) UpdateCategory(newCategory *models.Category, id string) (string, error) {
	var category *models.Category
	category, err := s.stores.Category.GetById(id)
	if err != nil {
		return "", err
	}
	category.Name = newCategory.Name
	category.Description = newCategory.Description

	err = s.stores.Category.Update(category)
	return category.ID, err
}

func (s *categoryService) DeleteCategory(id string) error {
	err := s.stores.Category.DeleteById(id)
	return err
}
