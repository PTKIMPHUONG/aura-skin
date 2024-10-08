package services

import (
	"auraskin/internal/models"
	"auraskin/internal/repositories"
	"errors"
)

type OrderServiceInterface interface {
	GetAllOrders() ([]models.Order, error)
	CreateOrder(order models.Order, userID string, productVariantID string) error
	UpdateOrder(orderID string, order models.Order) error
	CancelOrder(orderID string) error
	GetOrderByID(orderID string) (models.Order, error)
}

type orderService struct {
	orderRepo repositories.OrderRepository
}

func NewOrderService(orderRepo repositories.OrderRepository) OrderServiceInterface {
	return &orderService{orderRepo: orderRepo}
}

func (s *orderService) GetAllOrders() ([]models.Order, error) {
	return s.orderRepo.GetAllOrders()
}

func (s *orderService) CreateOrder(order models.Order, userID string, productVariantID string) error {
	if userID == "" || productVariantID == "" {
		return errors.New("userID and productVariantID are required")
	}
	return s.orderRepo.CreateOrder(order, userID, productVariantID)
}

func (s *orderService) UpdateOrder(orderID string, order models.Order) error {
	if orderID == "" {
		return errors.New("orderID is required")
	}
	return s.orderRepo.UpdateOrder(orderID, order)
}

func (s *orderService) CancelOrder(orderID string) error {
	if orderID == "" {
		return errors.New("orderID is required")
	}
	return s.orderRepo.CancelOrder(orderID)
}

func (s *orderService) GetOrderByID(orderID string) (models.Order, error) {
	if orderID == "" {
		return models.Order{}, errors.New("orderID is required")
	}
	return s.orderRepo.GetOrderByID(orderID)
}
