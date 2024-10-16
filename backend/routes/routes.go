package routes

import (
	"auraskin/internal/databases"

	"github.com/gofiber/fiber/v2"
)

var neo4jDB *databases.Neo4jDB

func SetupRoutes(app *fiber.App) {
	neo4jDB = databases.Instance()
	ProductRoutes(app)
	ProductVariantRoutes(app)
	CategoryRoutes(app)
	SupplierRoutes(app)
	ProductVariantRoutes(app)
	setupUserRoutes(app)
	setupOrderRoutes(app)
	SaleRoutes(app)
}
