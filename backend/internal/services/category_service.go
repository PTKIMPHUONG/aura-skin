package services

import (
	"auraskin/internal/models"
	"auraskin/internal/repositories"
)

type CategoryService interface {
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id string) (models.Category, error)
	GetProductsByCategoryID(categoryID string) ([]models.Product, error)  
	CreateCategory(category models.Category) error
	UpdateCategory(id string, category models.Category) error
	DeleteCategory(id string) error
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.GetAllCategories()
}

func (s *categoryService) GetCategoryByID(id string) (models.Category, error) {
	return s.repo.GetCategoryByID(id)
}

func (s *categoryService) GetProductsByCategoryID(categoryID string) ([]models.Product, error) {
	return s.repo.GetProductsByCategoryID(categoryID)
}

func (s *categoryService) CreateCategory(category models.Category) error {
	return s.repo.CreateCategory(category)
}

func (s *categoryService) UpdateCategory(id string, category models.Category) error {
	return s.repo.UpdateCategory(id, category)
}

func (s *categoryService) DeleteCategory(id string) error {
	return s.repo.DeleteCategory(id)
}
