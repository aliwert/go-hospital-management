package handlers

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"github.com/aliwert/go-hospital-management/internal/services"
	"github.com/aliwert/go-hospital-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type DepartmentHandler struct {
	departmentService *services.DepartmentService
}

func NewDepartmentHandler(departmentService *services.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{
		departmentService: departmentService,
	}
}

func (h *DepartmentHandler) CreateDepartment(c *fiber.Ctx) error {
	var req models.DepartmentCreateRequest
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

	department, err := h.departmentService.CreateDepartment(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(department)
}

func (h *DepartmentHandler) GetDepartments(c *fiber.Ctx) error {
	departments, err := h.departmentService.GetAllDepartments()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch departments",
		})
	}

	var response []models.DepartmentResponse
	for _, dept := range departments {
		response = append(response, models.DepartmentResponse{
			ID:          dept.ID,
			Name:        dept.Name,
			HeadDoctor:  dept.HeadDoctor.Name,
			Location:    dept.Location,
			Status:      dept.Status,
			StaffCount:  dept.StaffCount,
			OpenTime:    dept.OpenTime,
			CloseTime:   dept.CloseTime,
			PhoneNumber: dept.PhoneNumber,
		})
	}

	return c.JSON(response)
}

func (h *DepartmentHandler) GetDepartment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid department ID",
		})
	}

	department, err := h.departmentService.GetDepartmentById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Department not found",
		})
	}

	return c.JSON(models.DepartmentResponse{
		ID:          department.ID,
		Name:        department.Name,
		HeadDoctor:  department.HeadDoctor.Name,
		Location:    department.Location,
		Status:      department.Status,
		StaffCount:  department.StaffCount,
		OpenTime:    department.OpenTime,
		CloseTime:   department.CloseTime,
		PhoneNumber: department.PhoneNumber,
	})
}

func (h *DepartmentHandler) UpdateDepartment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid department ID",
		})
	}

	var req models.DepartmentUpdateRequest
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

	if err := h.departmentService.UpdateDepartment(uint(id), &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	department, _ := h.departmentService.GetDepartmentById(uint(id))
	return c.JSON(models.DepartmentResponse{
		ID:          department.ID,
		Name:        department.Name,
		HeadDoctor:  department.HeadDoctor.Name,
		Location:    department.Location,
		Status:      department.Status,
		StaffCount:  department.StaffCount,
		OpenTime:    department.OpenTime,
		CloseTime:   department.CloseTime,
		PhoneNumber: department.PhoneNumber,
	})
}

func (h *DepartmentHandler) DeleteDepartment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid department ID",
		})
	}

	if err := h.departmentService.DeleteDepartment(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
