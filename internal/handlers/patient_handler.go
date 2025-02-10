package handlers

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"github.com/aliwert/go-hospital-management/internal/services"
	"github.com/aliwert/go-hospital-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type PatientHandler struct {
	patientService *services.PatientService
}

func NewPatientHandler(patientService *services.PatientService) *PatientHandler {
	return &PatientHandler{
		patientService: patientService,
	}
}

func (h *PatientHandler) CreatePatient(c *fiber.Ctx) error {
	var req models.PatientCreateRequest
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

	patient, err := h.patientService.CreatePatient(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(patient)
}

func (h *PatientHandler) GetPatients(c *fiber.Ctx) error {
	patients, err := h.patientService.GetAllPatients()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch patients",
		})
	}

	var response []models.PatientResponse
	for _, patient := range patients {
		response = append(response, models.PatientResponse{
			ID:          patient.ID,
			Name:        patient.User.Name,
			Gender:      patient.Gender,
			BloodType:   patient.BloodType,
			PhoneNumber: patient.PhoneNumber,
			Insurance:   patient.Insurance,
			Status:      patient.Status,
			LastVisit:   patient.LastVisit,
			NextVisit:   patient.NextVisit,
		})
	}

	return c.JSON(response)
}

func (h *PatientHandler) GetPatient(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid patient ID",
		})
	}

	patient, err := h.patientService.GetPatientById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Patient not found",
		})
	}

	return c.JSON(models.PatientResponse{
		ID:          patient.ID,
		Name:        patient.User.Name,
		Gender:      patient.Gender,
		BloodType:   patient.BloodType,
		PhoneNumber: patient.PhoneNumber,
		Insurance:   patient.Insurance,
		Status:      patient.Status,
		LastVisit:   patient.LastVisit,
		NextVisit:   patient.NextVisit,
	})
}

func (h *PatientHandler) UpdatePatient(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid patient ID",
		})
	}

	patient, err := h.patientService.GetPatientById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Patient not found",
		})
	}

	if err := c.BodyParser(patient); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.patientService.UpdatePatient(uint(id), patient); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(patient)
}

func (h *PatientHandler) DeletePatient(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid patient ID",
		})
	}

	if err := h.patientService.DeletePatient(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
