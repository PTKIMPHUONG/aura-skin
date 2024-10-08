package routes

import (
	"auraskin/internal/controllers"
	"auraskin/internal/repositories"
	"auraskin/internal/services"

	"github.com/gofiber/fiber/v2"
)

func SupplierRoutes(app *fiber.App) {
	supplierRepo := repositories.NewSupplierRepository(neo4jDB)
	supplierService := services.NewSupplierService(supplierRepo)
	supplierController := controllers.NewSupplierController(supplierService)

	supplierGroup := app.Group("/suppliers")

	supplierGroup.Get("/", supplierController.GetAllSuppliers)
	supplierGroup.Get("/:id", supplierController.GetSupplierByID)
	supplierGroup.Post("/create", supplierController.CreateSupplier)
	supplierGroup.Put("/update/:id", supplierController.UpdateSupplier)
	supplierGroup.Delete("/delete/:id", supplierController.DeleteSupplier)
}