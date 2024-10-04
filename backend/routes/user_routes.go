package routes

import (
	"auraskin/internal/controllers"
	"auraskin/internal/middlewares"
	"auraskin/internal/repositories"
	"auraskin/internal/services"

	"github.com/gofiber/fiber/v2"
)

func setupUserRoutes(app *fiber.App) {
	repository := repositories.NewUserRepository(neo4jDB)
	service := services.NewUserService(repository)
	controller := controllers.NewUserController(service)

	userGroup := app.Group("/user")

	userGroup.Post("/register", controller.Register)
	userGroup.Post("/login", controller.Login)
	userGroup.Delete("/delete/:id", middlewares.AuthMiddleware(), controller.DeleteUser)
	userGroup.Put("/update", middlewares.AuthMiddleware(), controller.UpdateUser)
}
