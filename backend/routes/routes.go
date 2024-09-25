package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// userController := controllers.NewUserController(services.NewUserService(repositories.NewUserRepository(config.DB)))
	// productController := controllers.NewProductController(services.NewProductService(repositories.NewProductRepository(config.DB)))

	// app.Post("/register", userController.Register)
	// app.Post("/login", userController.Login)
	// app.Delete("/user/:email", middlewares.AuthMiddleware(), userController.DeleteUser)
	// app.Post("/add-product", middlewares.AuthMiddleware(), productController.AddProduct)
	// app.Get("/products", productController.GetProducts)
	// app.Put("/update-product", middlewares.AuthMiddleware(), productController.UpdateProduct)
	// app.Delete("/delete-product/:id", middlewares.AuthMiddleware(), productController.DeleteProduct)
}

//API request
//API gateway routes
//Controller
//Service - gọi Repository lấy data và xử lí nền(nếu có)
//Repository - Database
