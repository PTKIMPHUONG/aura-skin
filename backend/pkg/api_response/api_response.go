package utils

import (
    "github.com/gofiber/fiber/v2"
)

type SuccessResponse struct {
    Status  int         `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
    Status  int    `json:"status"`
    Message string `json:"message"`
    Error   string `json:"error,omitempty"`
}

func SendSuccessResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
    response := SuccessResponse{
        Status:  statusCode,
        Message: message,
        Data:    data,
    }
    return c.Status(statusCode).JSON(response)
}

func SendErrorResponse(c *fiber.Ctx, statusCode int, message string, err string) error {
    response := ErrorResponse{
        Status:  statusCode,
        Message: message,
        Error:   err,
    }
    return c.Status(statusCode).JSON(response)
}
