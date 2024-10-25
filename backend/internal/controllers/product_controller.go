package controllers

import (
	"auraskin/internal/models"
	"auraskin/internal/services"
	APIResponse "auraskin/pkg/api_response"
	"fmt"
	"net/url"
	"strconv"

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

func (pc *ProductController) UploadProductPicture(c *fiber.Ctx) error {
	productID := c.Params("product_id")

	fileHeader, err := c.FormFile("product_image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid file",
			Error:   err.Error(),
		})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to open file",
			Error:   err.Error(),
		})
	}
	defer file.Close()

	productPictureURL, err := pc.service.UploadProductPicture(productID, file, fileHeader)
	if err != nil {
		fmt.Println("Error in UploadProductPicture service:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to upload product picture",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Product picture uploaded successfully",
		Data:    fiber.Map{"product_picture_url": productPictureURL},
	})
}

func (pc *ProductController) GetProductByVariantID(c *fiber.Ctx) error {
	variantID := c.Params("variant_id")
	product, err := pc.service.GetProductByVariantID(variantID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusNotFound,
			Message: "Product not found for the specified variant",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Product retrieved successfully",
		Data:    product,
	})
}

func (pc *ProductController) GetProductByName(c *fiber.Ctx) error {
	productName := c.Query("product_name") 
	if productName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Product name is required",
		})
	}

	fmt.Println("GetProductByName called")
	decodedProductName, err := url.QueryUnescape(productName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Error decoding product name",
			Error:   err.Error(),
		})
	}

	products, err := pc.service.GetProductByName(decodedProductName)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusNotFound,
			Message: "Products not found",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Products retrieved successfully",
		Data:    products,
	})
}

func (pc *ProductController) FilterByPriceRange(c *fiber.Ctx) error {
    minPrice, err := strconv.ParseFloat(c.Query("min_price"), 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusBadRequest,
            Message: "Min price is not valid",
        })
    }

    maxPrice, err := strconv.ParseFloat(c.Query("max_price"), 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusBadRequest,
            Message: "Max price is not valid",
        })
    }

    products, err := pc.service.FilterByPriceRange(minPrice, maxPrice)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusNotFound,
            Message: "No products found in this price range",
            Error:   err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
        Status:  fiber.StatusOK,
        Message: "List of products by price range",
        Data:    products,
    })
}

func (pc *ProductController) SortByPrice(c *fiber.Ctx) error {
    order := c.Query("order", "asc") 
    
    products, err := pc.service.SortByPrice(order)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusInternalServerError,
            Message: "Unable to sort products",
            Error:   err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
        Status:  fiber.StatusOK,
        Message: "List of products sorted by price",
        Data:    products,
    })
}

func (pc *ProductController) SortByNewest(c *fiber.Ctx) error {
    products, err := pc.service.SortByNewest()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusInternalServerError,
            Message: "Unable to sort products",
            Error:   err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
        Status:  fiber.StatusOK,
        Message: "List of products in order of newest",
        Data:    products,
    })
}

func (pc *ProductController) GetProductsBySupplier(c *fiber.Ctx) error {
	supplierName := c.Query("supplier_name")
	if supplierName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Supplier name is required",
		})
	}

	products, err := pc.service.GetProductsBySupplier(supplierName)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusNotFound,
			Message: "Products from supplier not found",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Products retrieved successfully",
		Data:    products,
	})
}

func (pc *ProductController) FilterProducts(c *fiber.Ctx) error {
    categoryID := c.Query("category_id")
	supplierID := c.Query("supplier_id")
    minPrice, _ := strconv.ParseFloat(c.Query("min_price", "0"), 64)
    maxPrice, _ := strconv.ParseFloat(c.Query("max_price", "0"), 64)

    products, err := pc.service.FilterProducts(categoryID, supplierID, minPrice, maxPrice)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusNotFound,
            Message: "No products found with the applied filters",
            Error:   err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
        Status:  fiber.StatusOK,
        Message: "Filtered products retrieved successfully",
        Data:    products,
    })
}
