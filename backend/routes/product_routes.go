package routes

import (
	"auraskin/internal/controllers"
	"auraskin/internal/repositories"
	"auraskin/internal/services"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	storageRepository := repositories.NewStorageRepository()
	productRepo := repositories.NewProductRepository(neo4jDB, storageRepository)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	productGroup := app.Group("/products")

	productGroup.Get("/", productController.GetAllProducts)
	productGroup.Get("/:id", productController.GetProductByID)
	productGroup.Get("/:product_id/product-variants", productController.GetVariantsByProductID)
	productGroup.Get("/search/:product_name/product-variants", productController.GetVariantsByProductName)
	productGroup.Post("/create", productController.CreateProduct)
	productGroup.Put("/update/:id", productController.UpdateProduct)
	productGroup.Delete("/delete/:id", productController.DeleteProduct)
	productGroup.Post("/upload-product-picture/:product_id", productController.UploadProductPicture)
	productGroup.Get("/variant/:variant_id", productController.GetProductByVariantID)
}
