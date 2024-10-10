package routes

import (
	"auraskin/internal/controllers"
	"auraskin/internal/repositories"
	"auraskin/internal/services"

	"github.com/gofiber/fiber/v2"
)

func ProductVariantRoutes(app *fiber.App) {
	productVariantRepo := repositories.NewProductVariantRepository(neo4jDB)
	productVariantService := services.NewProductVariantService(productVariantRepo)
	productVariantController := controllers.NewProductVariantController(productVariantService)

	productVariantGroup := app.Group("/product-variants")

	productVariantGroup.Get("/", productVariantController.GetAllVariants)
	productVariantGroup.Get("/:id", productVariantController.GetVariantByID)
	productVariantGroup.Get("/search/:name", productVariantController.GetVariantByName)
	productVariantGroup.Post("/create", productVariantController.CreateVariant)
	productVariantGroup.Put("/update/:variant_id", productVariantController.UpdateVariant)
	productVariantGroup.Delete("/delete/:variant_id", productVariantController.DeleteVariant)
	productVariantGroup.Post("/upload-thumbnail/:variant_id", productVariantController.UploadThumbnail)
	productVariantGroup.Get("/suggest/user/:userID", productVariantController.GetSuggestVariantsForUser)
	productVariantGroup.Get("/suggest/variant/:id", productVariantController.GetSuggestVariantsForAVariant)
}
