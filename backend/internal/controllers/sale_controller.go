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
    // Lấy tham số `page` từ query, nếu không truyền sẽ mặc định là "1"
    page, err := strconv.Atoi(c.Query("page", "1"))
    if err != nil || page <= 0 {
        return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusBadRequest,
            Message: "Invalid page number",
            Error:   "Page number must be a positive integer",
        })
    }

    // Lấy tham số `pageSize` từ query, nếu không truyền sẽ mặc định là 4
    pageSize, err := strconv.Atoi(c.Query("pageSize", "4"))
    if err != nil || pageSize <= 0 {
        return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusBadRequest,
            Message: "Invalid page size",
            Error:   "Page size must be a positive integer",
        })
    }

    // Gọi service để lấy danh sách sales
    sales, err := sc.service.GetAllSales(page, pageSize)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusInternalServerError,
            Message: "Unable to fetch sales",
            Error:   err.Error(),
        })
    }

    // Trả về danh sách sales thành công
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
    // Lấy tham số dateStart từ query
    dateStart := c.Query("date_start")

    // Lấy tham số phân trang từ query
    page := c.QueryInt("page", 1)        
    pageSize := c.QueryInt("pageSize", 4) 

    // Gọi service để lấy danh sách sales theo ngày bắt đầu, có phân trang
    sales, err := sc.service.GetSalesByDateStart(dateStart, page, pageSize)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusInternalServerError,
            Message: "Unable to fetch sales by start date",
            Error:   err.Error(),
        })
    }

    // Trả về kết quả thành công
    return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
        Status:  fiber.StatusOK,
        Message: "Sales retrieved successfully",
        Data:    sales,
    })
}

func (sc *SaleController) GetSalesByDateEnd(c *fiber.Ctx) error {
    // Lấy tham số dateStart từ query
    dateEnd := c.Query("date_end")

    // Lấy tham số phân trang từ query
    page := c.QueryInt("page", 1)        
    pageSize := c.QueryInt("pageSize", 4) 

    // Gọi service để lấy danh sách sales theo ngày kết thúc, có phân trang
    sales, err := sc.service.GetSalesByDateEnd(dateEnd, page, pageSize)
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

func (sc *SaleController) GetExpiredSales(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid page number",
			Error:   "Page number must be a positive integer",
		})
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize", "10"))
	if err != nil || pageSize <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid page size",
			Error:   "Page size must be a positive integer",
		})
	}

	sales, err := sc.service.GetExpiredSales(page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to fetch expired sales",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Expired sales retrieved successfully",
		Data:    sales,
	})
}

func (sc *SaleController) SearchSalesByDescription(c *fiber.Ctx) error {
    description := c.Query("description")
    
    if description == "" {
        return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusBadRequest,
            Message: "Description parameter is required",
            Error:   "Missing 'query' parameter",
        })
    }

    page := c.QueryInt("page", 1)        
    pageSize := c.QueryInt("pageSize", 4) 

    sales, err := sc.service.SearchSalesByDescription(description, page, pageSize)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusInternalServerError,
            Message: "Unable to search sales by description",
            Error:   err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
        Status:  fiber.StatusOK,
        Message: "Sales retrieved successfully",
        Data:    sales,
    })
}

func (sc *SaleController) GetSalesByStatus(c *fiber.Ctx) error {
	isActive := c.Query("is_active") 
	var activeFlag bool
	if isActive == "true" {
		activeFlag = true
	} else if isActive == "false" {
		activeFlag = false
	}
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid page number",
			Error:   "Page number must be a positive integer",
		})
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize", "10"))
	if err != nil || pageSize <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid page size",
			Error:   "Page size must be a positive integer",
		})
	}
	sales, err := sc.service.GetSalesByStatus(activeFlag, page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Internal Server Error",
			Error:   "StatusInternalServerError",
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Sales by active retrieved successfully",
		Data:    sales,
	})
}