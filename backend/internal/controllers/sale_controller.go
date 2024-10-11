package controllers

import (
	"auraskin/internal/models"
	"auraskin/internal/services"
	APIResponse "auraskin/pkg/api_response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type SaleController struct {
	service services.SaleService
}

func NewSaleController(service services.SaleService) *SaleController {
	return &SaleController{service: service}
}

func (sc *SaleController) GetAllSales(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid page number",
			Error:   err.Error(),
		})
	}
	pageSize, err := strconv.Atoi(c.Query("pageSize", "10"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid page size",
			Error:   err.Error(),
		})
	}
	search := c.Query("search", "")

	sales, err := sc.service.GetAllSales(page, pageSize, search)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to fetch sales",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Sales retrieved successfully",
		Data:    sales,
	})
}

func (sc *SaleController) GetSaleByID(c *fiber.Ctx) error {
	id := c.Params("id")
	sale, err := sc.service.GetSaleByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusNotFound,
			Message: "Sale not found",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Sale retrieved successfully",
		Data:    sale,
	})
}

func (sc *SaleController) GetSalesByDateStart(c *fiber.Ctx) error {
	dateStart := c.Query("dateStart")
	sales, err := sc.service.GetSalesByDateStart(dateStart)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to fetch sales by start date",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Sales retrieved successfully",
		Data:    sales,
	})
}

func (sc *SaleController) GetSalesByDateEnd(c *fiber.Ctx) error {
	dateEnd := c.Query("dateEnd")
	sales, err := sc.service.GetSalesByDateEnd(dateEnd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to fetch sales by end date",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Sales retrieved successfully",
		Data:    sales,
	})
}

func (sc *SaleController) CreateSale(c *fiber.Ctx) error {
	var request struct {
		Sale      models.Sale `json:"sale"`
		VariantID string      `json:"variantID"`
	}

	// Parse body để lấy sale và variantID
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Cannot parse JSON",
			Error:   err.Error(),
		})
	}

	// Kiểm tra VariantID có tồn tại không
	if request.VariantID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Variant ID is required",
		})
	}

	// Gọi service để tạo sale kèm variantID
	if err := sc.service.CreateSale(request.Sale, request.VariantID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to create sale",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusCreated,
		Message: "Sale created successfully",
	})
}

func (sc *SaleController) UpdateSale(c *fiber.Ctx) error {
	id := c.Params("id")
	var sale models.Sale
	if err := c.BodyParser(&sale); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Cannot parse JSON",
			Error:   err.Error(),
		})
	}
	if err := sc.service.UpdateSale(id, sale); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to update sale",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Sale updated successfully",
	})
}

func (sc *SaleController) DeleteSale(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := sc.service.DeleteSale(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to delete sale",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Sale deleted successfully",
	})
}