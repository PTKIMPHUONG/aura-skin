package routes

import (
	"auraskin/internal/controllers"
	"auraskin/internal/repositories"
	"auraskin/internal/services"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App) {
	categoryRepo := repositories.NewCategoryRepository(neo4jDB)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryController := controllers.NewCategoryController(categoryService)

	categoryGroup := app.Group("/categories")

	categoryGroup.Get("/", categoryController.GetAllCategories)
	categoryGroup.Get("/:category_id", categoryController.GetCategoryByID)
	categoryGroup.Get("/:category_id/products", categoryController.GetProductsByCategoryID)
	categoryGroup.Post("/create", categoryController.CreateCategory)
	categoryGroup.Put("/update/:category_id", categoryController.UpdateCategory)
	categoryGroup.Delete("/delete/:category_id", categoryController.DeleteCategory)
}
