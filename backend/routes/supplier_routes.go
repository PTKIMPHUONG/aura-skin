package routes

import (
	"auraskin/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func SupplierRoutes(app *fiber.App, controller *controllers.SupplierController) {
	group := app.Group("/suppliers")

	group.Get("/", controller.GetAllSuppliers)
	group.Get("/:id", controller.GetSupplierByID)
	group.Post("/", controller.CreateSupplier)
	group.Put("/:id", controller.UpdateSupplier)
	group.Delete("/:id", controller.DeleteSupplier)
}