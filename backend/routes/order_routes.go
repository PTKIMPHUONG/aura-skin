package routes

import (
	"auraskin/internal/controllers"
	"auraskin/internal/repositories"
	"auraskin/internal/services"

	"github.com/gofiber/fiber/v2"
)

func setupOrderRoutes(app *fiber.App) {
	repository := repositories.NewOrderRepository(neo4jDB)
	service := services.NewOrderService(repository)
	controller := controllers.NewOrderController(service)

	orderGroup := app.Group("/order")

	orderGroup.Post("/create", controller.CreateOrder)
	orderGroup.Delete("/cancel/:id", controller.CancelOrder)
	orderGroup.Put("/update/:id", controller.UpdateOrder)
	orderGroup.Post("search/:id", controller.GetOrderByID)
}
