package handlers

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"github.com/aliwert/go-hospital-management/internal/services"
	"github.com/aliwert/go-hospital-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type AppointmentHandler struct {
	appointmentService *services.AppointmentService
}

func NewAppointmentHandler(appointmentService *services.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{
		appointmentService: appointmentService,
	}
}

func (h *AppointmentHandler) CreateAppointment(c *fiber.Ctx) error {
	var req models.AppointmentCreateRequest
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

	appointment, err := h.appointmentService.CreateAppointment(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(appointment)
}

func (h *AppointmentHandler) GetAppointments(c *fiber.Ctx) error {
	appointments, err := h.appointmentService.GetAllAppointments()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch appointments",
		})
	}

	var response []models.AppointmentResponse
	for _, apt := range appointments {
		response = append(response, models.AppointmentResponse{
			ID:              apt.ID,
			PatientName:     apt.Patient.Name,
			DoctorName:      apt.Doctor.Name,
			AppointmentDate: apt.AppointmentDate,
			Status:          apt.Status,
			Fee:             apt.Fee,
			PaymentStatus:   apt.PaymentStatus,
			CreatedAt:       apt.CreatedAt,
		})
	}

	return c.JSON(response)
}

func (h *AppointmentHandler) GetAppointment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid appointment ID",
		})
	}

	appointment, err := h.appointmentService.GetAppointmentById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Appointment not found",
		})
	}

	return c.JSON(models.AppointmentResponse{
		ID:              appointment.ID,
		PatientName:     appointment.Patient.Name,
		DoctorName:      appointment.Doctor.Name,
		AppointmentDate: appointment.AppointmentDate,
		Status:          appointment.Status,
		Fee:             appointment.Fee,
		PaymentStatus:   appointment.PaymentStatus,
		CreatedAt:       appointment.CreatedAt,
	})
}

func (h *AppointmentHandler) UpdateAppointment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid appointment ID",
		})
	}

	var req models.AppointmentUpdateRequest
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

	if err := h.appointmentService.UpdateAppointment(uint(id), &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *AppointmentHandler) DeleteAppointment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid appointment ID",
		})
	}

	if err := h.appointmentService.DeleteAppointment(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
