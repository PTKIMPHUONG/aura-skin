package routes

import (
	"auraskin/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func ProductVariantRoutes(app *fiber.App, productVariantController *controllers.ProductVariantController) {
	productVariantGroup := app.Group("/product-variants")

	productVariantGroup.Get("/", productVariantController.GetAllVariants)
	productVariantGroup.Get("/:id", productVariantController.GetVariantByID)
	productVariantGroup.Get("/:name", productVariantController.GetVariantByName)
	productVariantGroup.Post("/", productVariantController.CreateVariant)
	productVariantGroup.Put("/:variant_id", productVariantController.UpdateVariant)
	productVariantGroup.Delete("/:variant_id", productVariantController.DeleteVariant)
}
