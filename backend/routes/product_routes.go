package routes

import (
	"auraskin/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App, productController *controllers.ProductController) {
	productGroup := app.Group("/products")

	productGroup.Get("/", productController.GetAllProducts)
	productGroup.Get("/:id", productController.GetProductByID)
	productGroup.Get("/:product_id/product-variants", productController.GetVariantsByProductID)
	productGroup.Get("/:product_name/product-variants", productController.GetVariantsByProductName)
	productGroup.Post("/", productController.CreateProduct)
	productGroup.Put("/:id", productController.UpdateProduct)
	productGroup.Delete("/:id", productController.DeleteProduct)
}
