package handlers

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"github.com/aliwert/go-hospital-management/internal/services"
	"github.com/aliwert/go-hospital-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type SupplierHandler struct {
	supplierService *services.SupplierService
}

func NewSupplierHandler(supplierService *services.SupplierService) *SupplierHandler {
	return &SupplierHandler{
		supplierService: supplierService,
	}
}

func (h *SupplierHandler) CreateSupplier(c *fiber.Ctx) error {
	var req models.SupplierCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": utils.FormatValidationError(err),
		})
	}

	supplier, err := h.supplierService.CreateSupplier(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(supplier)
}

func (h *SupplierHandler) GetSuppliers(c *fiber.Ctx) error {
	suppliers, err := h.supplierService.GetAllSuppliers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch suppliers",
		})
	}

	return c.JSON(suppliers)
}

func (h *SupplierHandler) GetSupplier(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid supplier ID",
		})
	}

	supplier, err := h.supplierService.GetSupplierById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Supplier not found",
		})
	}

	return c.JSON(supplier)
}

func (h *SupplierHandler) UpdateSupplier(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid supplier ID",
		})
	}

	var req models.SupplierUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": utils.FormatValidationError(err),
		})
	}

	supplier, err := h.supplierService.GetSupplierById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Supplier not found",
		})
	}

	if req.Name != "" {
		supplier.Name = req.Name
	}
	if req.Email != "" {
		supplier.Email = req.Email
	}
	if req.Phone != "" {
		supplier.Phone = req.Phone
	}
	if req.Address != "" {
		supplier.Address = req.Address
	}
	if req.ContactPerson != "" {
		supplier.ContactPerson = req.ContactPerson
	}
	if req.ContactPhone != "" {
		supplier.ContactPhone = req.ContactPhone
	}
	if req.Status != "" {
		supplier.Status = req.Status
	}
	if req.PaymentTerms != "" {
		supplier.PaymentTerms = req.PaymentTerms
	}
	if req.DeliveryTerms != "" {
		supplier.DeliveryTerms = req.DeliveryTerms
	}

	if err := h.supplierService.UpdateSupplier(uint(id), supplier); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(supplier)
}

func (h *SupplierHandler) DeleteSupplier(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid supplier ID",
		})
	}

	if err := h.supplierService.DeleteSupplier(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
