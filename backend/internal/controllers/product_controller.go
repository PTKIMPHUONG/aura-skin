package controllers

import (
	"auraskin/internal/models"
	"auraskin/internal/services"
	APIResponse "auraskin/pkg/api_response"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	service services.ProductServiceInterface
}

func NewProductController(service services.ProductServiceInterface) *ProductController {
	return &ProductController{service}
}

func (pc *ProductController) GetAllProducts(c *fiber.Ctx) error {
	products, err := pc.service.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to fetch products",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Products retrieved successfully",
		Data:    products,
	})
}

func (pc *ProductController) GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	product, err := pc.service.GetProductByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusNotFound,
			Message: "Product not found",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Product retrieved successfully",
		Data:    product,
	})
}

func (pc *ProductController) GetVariantsByProductID(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	variants, err := pc.service.GetVariantsByProductID(productID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusNotFound,
			Message: "Variants not found for the specified product",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Variants retrieved successfully",
		Data:    variants,
	})
}

func (pc *ProductController) GetVariantsByProductName(c *fiber.Ctx) error {
	productName := c.Params("product_name")

	// Gọi service để lấy các variants dựa theo tên sản phẩm
	variants, err := pc.service.GetVariantsByProductName(productName)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusNotFound,
			Message: "Variants not found for the specified product name",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Variants retrieved successfully",
		Data:    variants,
	})
}

func (pc *ProductController) CreateProduct(c *fiber.Ctx) error {
    var request struct {
        Product    models.Product `json:"product"`
        CategoryID string         `json:"categoryID"`
        SupplierID string         `json:"supplierID"`
    }

    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusBadRequest,
            Message: "Cannot parse JSON",
            Error:   err.Error(),
        })
    }

    // Kiểm tra categoryID và supplierID có tồn tại không
    if request.CategoryID == "" {
        return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusBadRequest,
            Message: "Category ID is required",
        })
    }
    if request.SupplierID == "" {
        return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusBadRequest,
            Message: "Supplier ID is required",
        })
    }

    if err := pc.service.CreateProduct(request.Product, request.CategoryID, request.SupplierID); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusInternalServerError,
            Message: "Unable to create product",
            Error:   err.Error(),
        })
    }

    return c.Status(fiber.StatusCreated).JSON(APIResponse.SuccessResponse{
        Status:  fiber.StatusCreated,
        Message: "Product created successfully",
    })
}

func (pc *ProductController) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Cannot parse JSON",
			Error:   err.Error(),
		})
	}
	if err := pc.service.UpdateProduct(id, product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to update product",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Product updated successfully",
	})
}

func (pc *ProductController) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := pc.service.DeleteProduct(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to delete product",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Product deleted successfully",
	})
}