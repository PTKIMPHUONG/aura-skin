package services

import (
	"auraskin/internal/models"
	"auraskin/internal/repositories"
)

type SupplierService interface {
	GetAllSuppliers() ([]models.Supplier, error)
	GetSupplierByID(id string) (*models.Supplier, error)
	CreateSupplier(supplier models.Supplier) error
	UpdateSupplier(supplier models.Supplier) error
	DeleteSupplier(id string) error
}

type supplierService struct {
	repo repositories.SupplierRepository
}

func NewSupplierService(repo repositories.SupplierRepository) SupplierService {
	return &supplierService{repo: repo}
}

func (s *supplierService) GetAllSuppliers() ([]models.Supplier, error) {
	return s.repo.GetAllSuppliers()
}

func (s *supplierService) GetSupplierByID(id string) (*models.Supplier, error) {
	return s.repo.GetSupplierByID(id)
}

func (s *supplierService) CreateSupplier(supplier models.Supplier) error {
	return s.repo.CreateSupplier(supplier)
}

func (s *supplierService) UpdateSupplier(supplier models.Supplier) error {
	return s.repo.UpdateSupplier(supplier)
}

func (s *supplierService) DeleteSupplier(id string) error {
	return s.repo.DeleteSupplier(id)
}
