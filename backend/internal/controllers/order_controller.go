package controllers

import (
	"auraskin/internal/models"
	"auraskin/internal/services"
	APIResponse "auraskin/pkg/api_response"
	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
	service services.OrderServiceInterface
}

func NewOrderController(service services.OrderServiceInterface) *OrderController {
	return &OrderController{service}
}

func (oc *OrderController) GetAllOrders(c *fiber.Ctx) error {
	orders, err := oc.service.GetAllOrders()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to fetch orders",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Orders fetched successfully",
		Data:    orders,
	})
}

func (oc *OrderController) CreateOrder(c *fiber.Ctx) error {
	var request struct {
		Order           models.Order `json:"order"`
		UserID          string       `json:"userID"`
		ProductVariantID string      `json:"productVariantID"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Cannot parse JSON",
			Error:   err.Error(),
		})
	}

	if request.UserID == "" || request.ProductVariantID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "UserID and ProductVariantID are required",
		})
	}

	if err := oc.service.CreateOrder(request.Order, request.UserID, request.ProductVariantID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to create order",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusCreated,
		Message: "Order created successfully",
	})
}

func (oc *OrderController) UpdateOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Cannot parse JSON",
			Error:   err.Error(),
		})
	}

	if err := oc.service.UpdateOrder(orderID, order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to update order",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Order updated successfully",
	})
}

func (oc *OrderController) CancelOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")
	if err := oc.service.CancelOrder(orderID); err != nil {
		if err.Error() == "order does not exist" {
			return c.Status(fiber.StatusNotFound).JSON(APIResponse.ErrorResponse{
				Status:  fiber.StatusNotFound,
				Message: "Order not found",
				Error:   err.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to cancel order",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Order canceled successfully",
	})
}

func (oc *OrderController) GetOrderByID(c *fiber.Ctx) error {
	orderID := c.Params("id")

	order, err := oc.service.GetOrderByID(orderID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusNotFound,
			Message: "Order not found",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Order retrieved successfully",
		Data:    order,
	})
}