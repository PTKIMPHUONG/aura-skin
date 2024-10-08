package controllers

import (
	"auraskin/internal/models"
	"auraskin/internal/services"
	APIResponse "auraskin/pkg/api_response"
	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	service services.CategoryService
}

func NewCategoryController(service services.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (c *CategoryController) GetAllCategories(ctx *fiber.Ctx) error {
	categories, err := c.service.GetAllCategories()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return ctx.JSON(categories)
}

func (c *CategoryController) GetCategoryByID(ctx *fiber.Ctx) error {
	id := ctx.Params("category_id")
	category, err := c.service.GetCategoryByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}
	return ctx.JSON(category)
}

func (cc *CategoryController) GetProductsByCategoryID(c *fiber.Ctx) error {
	categoryID := c.Params("category_id")
	products, err := cc.service.GetProductsByCategoryID(categoryID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to fetch products by category",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Products retrieved successfully",
		Data:    products,
	})
}

func (c *CategoryController) CreateCategory(ctx *fiber.Ctx) error {
	var category models.Category
	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	if err := c.service.CreateCategory(category); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(fiber.StatusCreated).JSON(category)
}

func (c *CategoryController) UpdateCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("category_id")
	var category models.Category
	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	if err := c.service.UpdateCategory(id, category); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return ctx.JSON(category)
}

func (c *CategoryController) DeleteCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("category_id")
	if err := c.service.DeleteCategory(id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}