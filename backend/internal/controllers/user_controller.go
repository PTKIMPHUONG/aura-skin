package controllers

import (
	"auraskin/internal/models"
	"auraskin/internal/services"
	APIResponse "auraskin/pkg/api_response"
	"auraskin/pkg/jwt"
	"fmt"

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
}

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

func (pc *UserController) GetOrdersByUserID(c *fiber.Ctx) error {
	UserID := c.Params("id")
	orders, err := pc.service.GetOrdersByUserID(UserID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusNotFound,
			Message: "Orders not found for the specified user",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Orders retrieved successfully",
		Data:    orders,
	})
}
func (uc *UserController) UploadProfilePicture(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	fileHeader, err := c.FormFile("user_image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid file",
			Error:   err.Error(),
		})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to open file",
			Error:   err.Error(),
		})
	}
	defer file.Close()

	profilePictureURL, err := uc.service.UploadProfilePicture(userID, file, fileHeader)
	if err != nil {
		fmt.Println("Error in UploadProfilePicture service:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to upload profile picture",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Profile picture uploaded successfully",
		Data:    fiber.Map{"profile_picture_url": profilePictureURL},
	})
}

func (uc *UserController) GetByID(c *fiber.Ctx) error {
	userID := c.Params("id")
	user, err := uc.service.GetByID(userID)
	if err != nil {
		if err.Error() == "user not found" {
			return c.Status(fiber.StatusNotFound).JSON(APIResponse.ErrorResponse{
				Status:  fiber.StatusNotFound,
				Message: "User not found",
				Error:   "StatusNotFound",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Internal Server Error",
			Error:   "StatusInternalServerError",
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "User retrieved successfully",
		Data:    user,
	})
}

func (uc *UserController) GetUsersByName(c *fiber.Ctx) error {
	name := c.Query("username")
	users, err := uc.service.GetUsersByName(name)
	if err != nil {
		if err.Error() == "no users found with the name " + name {
			return c.Status(fiber.StatusNotFound).JSON(APIResponse.ErrorResponse{
				Status:  fiber.StatusNotFound,
				Message: "No users found with the given name",
				Error:   "StatusNotFound",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Internal Server Error",
			Error:   "StatusInternalServerError",
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Users retrieved successfully",
		Data:    users,
	})
}

func (uc *UserController) GetUserByEmail(c *fiber.Ctx) error {
    email := c.Query("email")  
    if email == "" {
        return c.Status(fiber.StatusBadRequest).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusBadRequest,
            Message: "Email query parameter is required",
            Error:   "StatusBadRequest",
        })
    }

    user, err := uc.service.GetUserByEmail(email)
    if err != nil {
        if err.Error() == "user with email " + email + " not found" {
            return c.Status(fiber.StatusNotFound).JSON(APIResponse.ErrorResponse{
                Status:  fiber.StatusNotFound,
                Message: "User not found",
                Error:   "StatusNotFound",
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusInternalServerError,
            Message: "Internal Server Error",
            Error:   "StatusInternalServerError",
        })
    }

    return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
        Status:  fiber.StatusOK,
        Message: "User retrieved successfully",
        Data:    user,
    })
}

func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := uc.service.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Internal Server Error",
			Error:   "StatusInternalServerError",
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Users retrieved successfully",
		Data:    users,
	})
}

func (uc *UserController) GetUserByRole(c *fiber.Ctx) error {
	isAdmin := c.Query("is_admin") 
	var adminFlag bool
	if isAdmin == "true" {
		adminFlag = true
	} else if isAdmin == "false" {
		adminFlag = false
	}

	users, err := uc.service.GetUserByRole(adminFlag)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Internal Server Error",
			Error:   "StatusInternalServerError",
		})
	}
	return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Users by role retrieved successfully",
		Data:    users,
	})
}

func (uc *UserController) GetProductVariantsByUserID(c *fiber.Ctx) error {
    userID := c.Params("id")

    productVariants, err := uc.service.GetProductVariantsByUserID(userID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(APIResponse.ErrorResponse{
            Status:  fiber.StatusInternalServerError,
            Message: "Internal Server Error",
            Error:   "StatusInternalServerError",
        })
    }

    return c.Status(fiber.StatusOK).JSON(APIResponse.SuccessResponse{
        Status:  fiber.StatusOK,
        Message: "Product variants retrieved successfully",
        Data:    productVariants,
    })
}
