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

	productGroup.Get("/search-by-name", productController.GetProductByName)
    productGroup.Get("/variant/:variant_id", productController.GetProductByVariantID)
    productGroup.Get("/search/:product_name/product-variants", productController.GetVariantsByProductName)
    productGroup.Post("/create", productController.CreateProduct)
    productGroup.Post("/upload-product-picture/:product_id", productController.UploadProductPicture)

    productGroup.Get("/:id", productController.GetProductByID)
    productGroup.Put("/update/:id", productController.UpdateProduct)
    productGroup.Delete("/delete/:id", productController.DeleteProduct)

    productGroup.Get("/", productController.GetAllProducts)
}
