package handlers

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"github.com/aliwert/go-hospital-management/internal/services"
	"github.com/aliwert/go-hospital-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type DoctorHandler struct {
	doctorService *services.DoctorService
}

func NewDoctorHandler(doctorService *services.DoctorService) *DoctorHandler {
	return &DoctorHandler{
		doctorService: doctorService,
	}
}

func (h *DoctorHandler) CreateDoctor(c *fiber.Ctx) error {
	var req models.DoctorCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": map[string]string{
				"body": "Invalid request body",
			},
		})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": utils.FormatValidationError(err),
		})
	}

	doctor, err := h.doctorService.CreateDoctor(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(models.DoctorResponse{
		ID:              doctor.ID,
		Name:            doctor.User.Name,
		Specialization:  doctor.Specialization,
		Department:      doctor.Department,
		Status:          doctor.Status,
		Availability:    doctor.Availability,
		ConsultationFee: doctor.ConsultationFee,
		Rating:          doctor.Rating,
		WorkingDays:     doctor.WorkingDays,
		WorkingHours:    doctor.WorkingHours,
	})
}
func (h *DoctorHandler) UpdateDoctor(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid doctor ID",
		})
	}

	var req models.DoctorUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": map[string]string{
				"body": "Invalid request body",
			},
		})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": utils.FormatValidationError(err),
		})
	}

	if err := h.doctorService.UpdateDoctor(uint(id), &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	doctor, err := h.doctorService.GetDoctorById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Doctor not found",
		})
	}

	return c.JSON(models.DoctorResponse{
		ID:              doctor.ID,
		Name:            doctor.User.Name,
		Specialization:  doctor.Specialization,
		Department:      doctor.Department,
		Status:          doctor.Status,
		Availability:    doctor.Availability,
		ConsultationFee: doctor.ConsultationFee,
		Rating:          doctor.Rating,
		WorkingDays:     doctor.WorkingDays,
		WorkingHours:    doctor.WorkingHours,
	})
}

func (h *DoctorHandler) DeleteDoctor(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid doctor ID",
		})
	}

	if err := h.doctorService.DeleteDoctor(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (h *DoctorHandler) GetDoctors(c *fiber.Ctx) error {
	doctors, err := h.doctorService.GetAllDoctors()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch doctors",
		})
	}

	var response []models.DoctorResponse
	for _, doctor := range doctors {
		response = append(response, models.DoctorResponse{
			ID:              doctor.ID,
			Name:            doctor.User.Name,
			Specialization:  doctor.Specialization,
			Department:      doctor.Department,
			Status:          doctor.Status,
			Availability:    doctor.Availability,
			ConsultationFee: doctor.ConsultationFee,
			Rating:          doctor.Rating,
			WorkingDays:     doctor.WorkingDays,
			WorkingHours:    doctor.WorkingHours,
		})
	}

	return c.JSON(response)
}

func (h *DoctorHandler) GetDoctor(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid doctor ID",
		})
	}

	doctor, err := h.doctorService.GetDoctorById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Doctor not found",
		})
	}

	return c.JSON(models.DoctorResponse{
		ID:              doctor.ID,
		Name:            doctor.User.Name,
		Specialization:  doctor.Specialization,
		Department:      doctor.Department,
		Status:          doctor.Status,
		Availability:    doctor.Availability,
		ConsultationFee: doctor.ConsultationFee,
		Rating:          doctor.Rating,
		WorkingDays:     doctor.WorkingDays,
		WorkingHours:    doctor.WorkingHours,
	})
}
