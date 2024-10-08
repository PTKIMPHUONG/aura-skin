package routes

import (
	"auraskin/internal/controllers"
	"auraskin/internal/databases"
	"auraskin/internal/repositories"
	"auraskin/internal/services"

	"github.com/gofiber/fiber/v2"
)

var neo4jDB *databases.Neo4jDB

func SetupRoutes(app *fiber.App) {
	neo4jDB = databases.Instance()

	productRepo := repositories.NewProductRepository(neo4jDB)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)
	ProductRoutes(app, productController)

	productVariantRepo := repositories.NewProductVariantRepository(neo4jDB)
	productVariantService := services.NewProductVariantService(productVariantRepo)
	productVariantController := controllers.NewProductVariantController(productVariantService)
	ProductVariantRoutes(app, productVariantController)

	categoryRepo := repositories.NewCategoryRepository(neo4jDB)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryController := controllers.NewCategoryController(categoryService)
	CategoryRoutes(app, categoryController)

	supplierRepo := repositories.NewSupplierRepository(neo4jDB)
	supplierService := services.NewSupplierService(supplierRepo)
	supplierController := controllers.NewSupplierController(supplierService)
	SupplierRoutes(app, supplierController)

	productVariantRepo = repositories.NewProductVariantRepository(neo4jDB)
	productVariantService = services.NewProductVariantService(productVariantRepo)
	productVariantController = controllers.NewProductVariantController(productVariantService)
	ProductVariantRoutes(app, productVariantController)
}
