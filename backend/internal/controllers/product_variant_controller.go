package controllers

import (
	"auraskin/internal/models"
	"auraskin/internal/services"
	APIResponse "auraskin/pkg/api_response"
	"net/url"
	"github.com/gofiber/fiber/v2"
)

type ProductVariantController struct {
	service services.ProductVariantService
}

func NewProductVariantController(service services.ProductVariantService) *ProductVariantController {
	return &ProductVariantController{service}
}

func (vc *ProductVariantController) GetAllVariants(c *fiber.Ctx) error {
	variants, err := vc.service.GetAllVariants()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to fetch variants",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Variants retrieved successfully",
		Data:    variants,
	})
}

func (vc *ProductVariantController) GetVariantByID(c *fiber.Ctx) error {
	id := c.Params("id")
	variant, err := vc.service.GetVariantByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusNotFound,
			Message: "Variant not found",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Variant retrieved successfully",
		Data:    variant,
	})
}

func (pc *ProductVariantController) GetVariantByName(c *fiber.Ctx) error {
    name := c.Params("name")
    decodedName, err := url.QueryUnescape(name)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusBadRequest,
            Message: "Invalid variant name format",
            Error:   err.Error(),
        })
    }

    variant, err := pc.service.GetVariantByName(decodedName)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusInternalServerError,
            Message: "Unable to fetch variant by name",
            Error:   err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
        Status:  fiber.StatusOK,
        Data:    variant,
    })
}

func (vc *ProductVariantController) CreateVariant(c *fiber.Ctx) error {
    // Struct để chứa dữ liệu từ request body
    var request struct {
        Variant   models.ProductVariant `json:"variant"`
        ProductID string                `json:"productID"`
    }

    // Parse body để lấy variant và productID
    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusBadRequest,
            Message: "Cannot parse JSON",
            Error:   err.Error(),
        })
    }

    // Kiểm tra ProductID có tồn tại không
    if request.ProductID == "" {
        return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusBadRequest,
            Message: "Product ID is required",
        })
    }

    // Gọi service để tạo variant kèm productID
    if err := vc.service.CreateVariant(request.Variant, request.ProductID); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusInternalServerError,
            Message: "Unable to create variant",
            Error:   err.Error(),
        })
    }

    return c.Status(fiber.StatusCreated).JSON(APIResponse.SuccessResponse{
        Status:  fiber.StatusCreated,
        Message: "Variant created successfully",
    })
}

func (vc *ProductVariantController) UpdateVariant(c *fiber.Ctx) error {
	id := c.Params("id")
	var variant models.ProductVariant
	if err := c.BodyParser(&variant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Cannot parse JSON",
			Error:   err.Error(),
		})
	}
	if err := vc.service.UpdateVariant(id, variant); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to update variant",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Variant updated successfully",
	})
}

func (vc *ProductVariantController) DeleteVariant(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := vc.service.DeleteVariant(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to delete variant",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Variant deleted successfully",
	})
}
func (vc *ProductVariantController) UploadThumbnail(c *fiber.Ctx) error {
	variantID := c.Params("variant_id")

	fileHeader, err := c.FormFile("thumbnail")
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

	thumbnailURL, err := vc.service.UploadThumbnailAndSetURL(variantID, file, fileHeader)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to upload and update thumbnail",
			Error:   err.Error(),
		})
	}

	return c.JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Thumbnail uploaded successfully",
		Data:    fiber.Map{"thumbnail_url": thumbnailURL},
	})
}