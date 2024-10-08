package controllers

import (
	"auraskin/internal/models"
	"auraskin/internal/services"
	APIResponse "auraskin/pkg/api_response"
	"auraskin/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service}
}

// Register
func (uc *UserController) Register(c *fiber.Ctx) error {
	// Parse request body to User model
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	// Call the service to register the user
	err := uc.service.Register(&user)
	if err != nil {
		// Check if the error is due to the email already existing
		if err.Error() == "email already exists" {
			return c.Status(fiber.StatusConflict).JSON(APIResponse.ErrorResponse{
				Status:  fiber.StatusConflict,
				Message: "Email already exists",
				Error:   "Conflict",
			})
		}

		// Return a general internal server error for other cases
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Internal Server Error",
			Error:   "StatusInternalServerError",
		})
	}

	// Return success response on successful registration
	return c.Status(fiber.StatusCreated).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusCreated,
		Message: "Register successful",
		Data:    user,
	})
}

// Login
func (uc *UserController) Login(c *fiber.Ctx) error {
	var req models.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	//Fetch user info
	existingUser, err := uc.service.Login(req.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusUnauthorized,
			Message: "Invalid email",
			Error:   "StatusUnauthorized",
		})
	}
	//Check password
	err = uc.service.ComparePassword(existingUser.Password, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusUnauthorized,
			Message: "Invalid password",
			Error:   "StatusUnauthorized",
		})
	}
	//JWT token
	token, err := jwt.GenerateToken(*existingUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Could not create token",
			Error:   "StatusInternalServerError",
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Login succesful",
		Data: fiber.Map{
			"token": token,
			"data":  existingUser,
		},
	})
}

// Delete
func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
    if !ok || userID == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusUnauthorized,
            Message: "Unauthorized",
            Error:   "StatusUnauthorized",
        })
    }
	isAdmin := c.Locals("isAdmin").(bool)
	id := c.Params("id")

	if id != userID && !isAdmin {
        return c.Status(fiber.StatusUnauthorized).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusUnauthorized,
            Message: "Unauthorized",
            Error:   "StatusUnauthorized",
        })
    }

	err := uc.service.DeleteAccount(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusNotFound,
			Message: "User not found",
			Error:   "StatusNotFound",
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "User deleted succesfully",
		Data:    nil,
	})


// Update
func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
    var updatedData models.User
    if err := c.BodyParser(&updatedData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

	userID, ok := c.Locals("userID").(string)
    if !ok || userID == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusUnauthorized,
            Message: "Unauthorized",
            Error:   "StatusUnauthorized",
        })
    }
	
	updatedData.ID = userID
    err := uc.service.Update(&updatedData)
    if err != nil {
        if err.Error() == "user not found" {
            return c.Status(fiber.StatusNotFound).JSON(APIResponse.ErrorResponse{
                Status:  fiber.StatusNotFound,
                Message: "User not found",
                Error:   "StatusNotFound",
            })
        } else if err.Error() == "you do not have permission to update this user" {
            return c.Status(fiber.StatusForbidden).JSON(APIResponse.ErrorResponse{
                Status:  fiber.StatusForbidden,
                Message: "You do not have permission to update this user",
                Error:   "StatusForbidden",
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusInternalServerError,
            Message: "Failed to update user information",
            Error:   "StatusInternalServerError",
        })
    }

    return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
        Status:  fiber.StatusOK,
        Message: "User updated successfully",
        Data:    updatedData,
    })
}

