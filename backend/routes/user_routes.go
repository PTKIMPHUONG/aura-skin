package routes

import (
	"auraskin/internal/controllers"
	"auraskin/internal/middlewares"
	"auraskin/internal/repositories"
	"auraskin/internal/services"

	"github.com/gofiber/fiber/v2"
)

func setupUserRoutes(app *fiber.App) {
	storageRepository := repositories.NewStorageRepository()
	repository := repositories.NewUserRepository(neo4jDB, storageRepository)
	service := services.NewUserService(repository)
	controller := controllers.NewUserController(service)

	userGroup := app.Group("/user")

	userGroup.Post("/register", controller.Register)
	userGroup.Post("/login", controller.Login)
	userGroup.Delete("/delete/:id", middlewares.AuthMiddleware(), controller.DeleteUser)
	userGroup.Put("/update", middlewares.AuthMiddleware(), controller.UpdateUser)
	userGroup.Get("/:id/order-history", controller.GetOrdersByUserID)
	userGroup.Post("/upload-profile-picture/:user_id", controller.UploadProfilePicture)
	userGroup.Get("/:id", controller.GetByID)
	userGroup.Get("/users/username", controller.GetUsersByName)
	userGroup.Get("/search/email", controller.GetUserByEmail)
	userGroup.Get("/", controller.GetAllUsers) 
	userGroup.Get("/users/admin", controller.GetUserByRole)   
	userGroup.Get("/:id/product-variants", controller.GetProductVariantsByUserID)
}
