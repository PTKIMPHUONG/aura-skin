package services

import (
	"auraskin/internal/models"
	"auraskin/internal/repositories"
)

type SaleService interface {
	GetAllSales(page int, pageSize int, search string) ([]models.Sale, error)
	GetSaleByID(id string) (models.Sale, error)
	GetSalesByDateStart(dateStart string) ([]models.Sale, error)
	GetSalesByDateEnd(dateEnd string) ([]models.Sale, error)
	CreateSale(sale models.Sale, variantID string) error
	UpdateSale(id string, sale models.Sale) error
	DeleteSale(id string) error
}

type saleService struct {
	repo repositories.SaleRepository
}

func NewSaleService(repo repositories.SaleRepository) SaleService {
	return &saleService{repo: repo}
}

func (s *saleService) GetAllSales(page int, pageSize int, search string) ([]models.Sale, error) {
	return s.repo.GetAllSales(page, pageSize, search)
}

func (s *saleService) GetSaleByID(id string) (models.Sale, error) {
	return s.repo.GetSaleByID(id)
}

func (s *saleService) GetSalesByDateStart(dateStart string) ([]models.Sale, error) {
	return s.repo.GetSalesByDateStart(dateStart)
}

func (s *saleService) GetSalesByDateEnd(dateEnd string) ([]models.Sale, error) {
	return s.repo.GetSalesByDateEnd(dateEnd)
}

func (s *saleService) CreateSale(sale models.Sale, variantID string) error {
	return s.repo.CreateSale(sale, variantID)
}

func (s *saleService) UpdateSale(id string, sale models.Sale) error {
	return s.repo.UpdateSale(id, sale)
}

func (s *saleService) DeleteSale(id string) error {
	return s.repo.DeleteSale(id)
}
