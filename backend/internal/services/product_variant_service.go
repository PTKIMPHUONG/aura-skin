package services

import (
	"auraskin/internal/models"
	"auraskin/internal/repositories"
)

type ProductVariantService interface {
	GetAllVariants() ([]models.ProductVariant, error)
	GetVariantByID(id string) (models.ProductVariant, error)
	GetVariantByName(name string) (models.ProductVariant, error)
	CreateVariant(variant models.ProductVariant, productID string) error
	UpdateVariant(id string, variant models.ProductVariant) error
	DeleteVariant(id string) error
}

type productVariantService struct {
	repo repositories.ProductVariantRepository
}

func NewProductVariantService(repo repositories.ProductVariantRepository) ProductVariantService {
	return &productVariantService{repo: repo}
}

func (s *productVariantService) GetAllVariants() ([]models.ProductVariant, error) {
	return s.repo.GetAllVariants()
}

func (s *productVariantService) GetVariantByID(id string) (models.ProductVariant, error) {
	return s.repo.GetVariantByID(id)
}

func (s *productVariantService) GetVariantByName(name string) (models.ProductVariant, error) {
	return s.repo.GetVariantByName(name)
}

func (s *productVariantService) CreateVariant(variant models.ProductVariant, productID string) error {
	return s.repo.CreateVariant(variant, productID)
}

func (s *productVariantService) UpdateVariant(id string, variant models.ProductVariant) error {
	return s.repo.UpdateVariant(id, variant)
}

func (s *productVariantService) DeleteVariant(id string) error {
	return s.repo.DeleteVariant(id)
}