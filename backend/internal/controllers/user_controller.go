package controllers

// import (
// 	"firstproject/models"
// 	"firstproject/services"
// 	"firstproject/utils"

// 	"github.com/gofiber/fiber/v2"
// 	"golang.org/x/crypto/bcrypt"
// )

// type UserController struct {
// 	service services.UserService
// }

// func NewUserController(service services.UserService) *UserController {
// 	return &UserController{service}
// }

// func (uc *UserController) Register(c *fiber.Ctx) error {
// 	var user models.User
// 	if err := c.BodyParser(&user); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
// 	}
// 	if err := uc.service.Register(&user); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse{
// 			Status:  fiber.StatusInternalServerError,
// 			Message: "User already exists",
// 			Error:   "StatusInternalServerError",
// 		})

// 	}
// 	return c.Status(fiber.StatusCreated).JSON(utils.SuccessResponse{
// 		Status:  fiber.StatusCreated,
// 		Message: "Register successful",
// 		Data:    user,
// 	})
// }

// func (uc *UserController) Login(c *fiber.Ctx) error {
// 	var user models.User
// 	if err := c.BodyParser(&user); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
// 	}

// 	existingUser, err := uc.service.FindUserByEmail(user.Email)
// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(utils.ErrorResponse{
// 			Status:  fiber.StatusUnauthorized,
// 			Message: "Invalid email",
// 			Error:   "StatusUnauthorized",
// 		})
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(utils.ErrorResponse{
// 			Status:  fiber.StatusUnauthorized,
// 			Message: "Invalid password",
// 			Error:   "StatusUnauthorized",
// 		})
// 	}

// 	token, err := utils.GenerateToken(*existingUser)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse{
// 			Status:  fiber.StatusInternalServerError,
// 			Message: "Could not create token",
// 			Error:   "StatusInternalServerError",
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse{
// 		Status:  fiber.StatusOK,
// 		Message: "Login successful",
// 		Data: fiber.Map{
// 			"token": token,
// 			"data":  existingUser,
// 		},
// 	})
// }

// func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
// 	email := c.Params("email")
//     if email == "" {
// 		return c.Status(fiber.StatusNotFound).JSON(utils.ErrorResponse{
// 			Status:  fiber.StatusNotFound,
// 			Message: "User not found",
// 			Error:   "StatusNotFound",
// 		})
// 	}

// 	err := uc.service.DeleteUser(email)
// 	if err != nil {
// 		return c.Status(fiber.StatusNotFound).JSON(utils.ErrorResponse{
// 			Status:  fiber.StatusNotFound,
// 			Message: "User not found",
// 			Error:   "StatusNotFound",
// 		})
// 	}
// 	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse{
// 		Status:  fiber.StatusOK,
// 		Message: "User deleted successfully",
// 		Data:    nil,
// 	})
// }