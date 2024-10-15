package services

import (
	"auraskin/internal/models"
	"auraskin/internal/repositories"
	// "mime/multipart"
)

type ProductServiceInterface interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id string) (models.Product, error)
	GetVariantsByProductID(productID string) ([]models.ProductVariant, error)
	GetVariantsByProductName(productName string) ([]models.ProductVariant, error)
	CreateProduct(product models.Product, categoryID string, supplierID string) error
	UpdateProduct(id string, product models.Product) error
	DeleteProduct(id string) error
}

type ProductService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductServiceInterface {
	return &ProductService{repo}
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *ProductService) GetProductByID(id string) (models.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *ProductService) CreateProduct(product models.Product, categoryID string, supplierID string) error {
	return s.repo.CreateProduct(product, categoryID, supplierID)
}

func (s *ProductService) UpdateProduct(id string, product models.Product) error {
	return s.repo.UpdateProduct(id, product)
}

func (s *ProductService) DeleteProduct(id string) error {
	return s.repo.DeleteProduct(id)
}

func (s *ProductService) GetVariantsByProductID(productID string) ([]models.ProductVariant, error) {
	return s.repo.GetVariantsByProductID(productID)
}

func (s *ProductService) GetVariantsByProductName(productName string) ([]models.ProductVariant, error) {
	return s.repo.GetVariantsByProductName(productName)
}
