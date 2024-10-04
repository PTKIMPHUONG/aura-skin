package middlewares

import (
	config "auraskin/internal/configs/dev"
	APIResponse "auraskin/pkg/api_response"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		log.Println("Authorization Header:", token)

		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid token"})
		}
		token = strings.TrimPrefix(token, "Bearer ")

		claims := jwt.MapClaims{}
		cfg, err := config.Instance()
		if err != nil {
			log.Println("Error loading config:", err)
			return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{"error": "Can't load credentials"})
		}

		tkn, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(cfg.GetSecretKey()), nil // Sử dụng SecretKey từ cấu hình
		})

		if err != nil {
			log.Println("Error parsing token:", err)
			return c.Status(fiber.StatusUnauthorized).JSON(APIResponse.ErrorResponse{
				Status:  fiber.StatusUnauthorized,
				Message: "Invalid token: Token has expired",
				Error:   "StatusUnauthorized",
			})
		}

		if !tkn.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(APIResponse.ErrorResponse{
				Status:  fiber.StatusUnauthorized,
				Message: "Invalid token",
				Error:   "StatusUnauthorized",
			})
		}

		userID, okUserID := claims["userID"].(string)
		isAdmin, okIsAdmin := claims["isAdmin"].(bool)

		if !okUserID || userID == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(APIResponse.ErrorResponse{
				Status:  fiber.StatusUnauthorized,
				Message: "Unauthorized",
				Error:   "StatusUnauthorized",
			})
		}
		if !okIsAdmin {
			isAdmin = false
		}

		c.Locals("userID", userID)
		c.Locals("isAdmin", isAdmin)

		return c.Next()
	}
}
