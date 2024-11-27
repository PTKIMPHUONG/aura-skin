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

	// Authentication routes
	userGroup.Post("/register", controller.Register)
	userGroup.Post("/login", controller.Login)

	// User management routes (Admin)
	userGroup.Delete("/delete/:id", middlewares.AuthMiddleware(), controller.DeleteUser)
	userGroup.Put("/update", middlewares.AuthMiddleware(), controller.UpdateUser)

	// User information routes
	userGroup.Get("/:id", controller.GetByID)
	userGroup.Get("/", controller.GetAllUsers)
	userGroup.Get("/users/username", controller.GetUsersByName)
	userGroup.Get("/search/email", controller.GetUserByEmail)
	userGroup.Get("/users/admin", controller.GetUserByRole)
	userGroup.Get("/:id/order-history", controller.GetOrdersByUserID)
	userGroup.Get("/:id/product-variants", controller.GetProductVariantsByUserID)

	// Profile picture routes
	userGroup.Post("/upload-profile-picture/:user_id", controller.UploadProfilePicture)

	// Wishlist routes
	userGroup.Post("/:user_id/wishlist/:variant_id", controller.AddToWishlist)  
	userGroup.Delete("/:user_id/wishlist/:variant_id", controller.RemoveFromWishlist) 
	userGroup.Get("/:user_id/wishlist", controller.GetUserWishlist)

	// Cart routes
	userGroup.Post("/:user_id/cart", controller.AddToCart) 
	userGroup.Delete("/:user_id/cart/:variant_id", controller.RemoveFromCart)
	userGroup.Get("/:user_id/cart", controller.GetUserCart) 
}
