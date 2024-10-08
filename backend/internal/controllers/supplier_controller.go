package controllers

import (
	"auraskin/internal/models"
	"auraskin/internal/services"
	"github.com/gofiber/fiber/v2"
)

type SupplierController struct {
	service services.SupplierService
}

func NewSupplierController(service services.SupplierService) *SupplierController {
	return &SupplierController{service: service}
}

func (sc *SupplierController) GetAllSuppliers(c *fiber.Ctx) error {
	suppliers, err := sc.service.GetAllSuppliers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Unable to retrieve suppliers",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"data":    suppliers,
	})
}

func (sc *SupplierController) GetSupplierByID(c *fiber.Ctx) error {
	id := c.Params("id")
	supplier, err := sc.service.GetSupplierByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Supplier not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"data":    supplier,
	})
}

func (sc *SupplierController) CreateSupplier(c *fiber.Ctx) error {
	var supplier models.Supplier
	if err := c.BodyParser(&supplier); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	if err := sc.service.CreateSupplier(supplier); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Unable to create supplier",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "Supplier created successfully",
	})
}

func (sc *SupplierController) UpdateSupplier(c *fiber.Ctx) error {
	id := c.Params("id")
	var supplier models.Supplier
	if err := c.BodyParser(&supplier); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}
	supplier.SupplierID = id

	if err := sc.service.UpdateSupplier(supplier); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Unable to update supplier",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Supplier updated successfully",
	})
}

func (sc *SupplierController) DeleteSupplier(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := sc.service.DeleteSupplier(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Unable to delete supplier",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusNoContent).JSON(nil)
}
