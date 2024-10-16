package services

import (
	"auraskin/internal/models"
	"auraskin/internal/repositories"
)

type SaleService interface {
	GetAllSales(page int, pageSize int) ([]map[string]interface{}, error)
	GetSaleByID(id string) (map[string]interface{}, error)
	GetSalesByDateStart(dateStart string, page int, pageSize int) ([]map[string]interface{}, error)
	GetSalesByDateEnd(dateEnd string, page int, pageSize int) ([]map[string]interface{}, error)
	CreateSale(sale models.Sale, variantID string) error
	UpdateSale(id string, sale models.Sale) error
	DeleteSale(id string) error
	GetExpiredSales(page int, pageSize int) ([]map[string]interface{}, error)
	SearchSalesByDescription(description string, page int, pageSize int) ([]map[string]interface{}, error)
	GetSalesByStatus(isActive bool, page int, pageSize int) ([]map[string]interface{}, error)
}

type saleService struct {
	repo repositories.SaleRepository
}

func NewSaleService(repo repositories.SaleRepository) SaleService {
	return &saleService{repo: repo}
}

func (s *saleService) GetAllSales(page int, pageSize int) ([]map[string]interface{}, error) {
	return s.repo.GetAllSales(page, pageSize)
}

func (s *saleService) GetSaleByID(id string) (map[string]interface{}, error) {
	return s.repo.GetSaleByID(id)
}

func (s *saleService) GetSalesByDateStart(dateStart string, page int, pageSize int) ([]map[string]interface{}, error) {
    return s.repo.GetSalesByDateStart(dateStart, page, pageSize)
}

func (s *saleService) GetSalesByDateEnd(dateEnd string, page int, pageSize int) ([]map[string]interface{}, error) {
    return s.repo.GetSalesByDateEnd(dateEnd, page, pageSize)
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

func (s *saleService) GetSalesByStatus(isActive bool, page int, pageSize int) ([]map[string]interface{}, error) {
	return s.repo.GetSalesByStatus(isActive, page, pageSize)
}

func (s *saleService) GetExpiredSales(page int, pageSize int) ([]map[string]interface{}, error) {
	return s.repo.GetExpiredSales(page, pageSize)
}

func (s *saleService) SearchSalesByDescription(description string, page int, pageSize int) ([]map[string]interface{}, error) {
	return s.repo.SearchSalesByDescription(description, page, pageSize)
}