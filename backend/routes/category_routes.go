package routes

import (
	"auraskin/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App, categoryController *controllers.CategoryController) {
	categoryGroup := app.Group("/categories")

	categoryGroup.Get("/", categoryController.GetAllCategories)
	categoryGroup.Get("/:category_id", categoryController.GetCategoryByID)
	categoryGroup.Get("/:category_id/products", categoryController.GetProductsByCategoryID)
	categoryGroup.Post("/", categoryController.CreateCategory)
	categoryGroup.Put("/:category_id", categoryController.UpdateCategory)
	categoryGroup.Delete("/:category_id", categoryController.DeleteCategory)
}
