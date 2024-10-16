package routes

import (
	"auraskin/internal/controllers"
	"auraskin/internal/repositories"
	"auraskin/internal/services"

	"github.com/gofiber/fiber/v2"
)

func SaleRoutes(app *fiber.App) {
	saleRepo := repositories.NewSaleRepository(neo4jDB)
	saleService := services.NewSaleService(saleRepo)
	saleController := controllers.NewSaleController(saleService)

	saleGroup := app.Group("/sales")

	saleGroup.Get("/", saleController.GetAllSales)                   
	saleGroup.Get("/:id", saleController.GetSaleByID)                
	saleGroup.Get("/search/start-date", saleController.GetSalesByDateStart)  
	saleGroup.Get("/search/end-date", saleController.GetSalesByDateEnd)      
	saleGroup.Post("/create", saleController.CreateSale)              
	saleGroup.Put("/update/:id", saleController.UpdateSale)           
	saleGroup.Delete("/delete/:id", saleController.DeleteSale)      
	saleGroup.Get("/filter-by/active-status", saleController.GetSalesByStatus)               
	saleGroup.Get("/expired-sales", saleController.GetExpiredSales)              
	saleGroup.Get("/search-by-description", saleController.SearchSalesByDescription)
}
