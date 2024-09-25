package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func ExceptionMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next() // Tiến hành xử lý request
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return nil
	}
}
