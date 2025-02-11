package handlers

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"github.com/aliwert/go-hospital-management/internal/services"
	"github.com/aliwert/go-hospital-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type PrescriptionHandler struct {
	prescriptionService *services.PrescriptionService
}

func NewPrescriptionHandler(prescriptionService *services.PrescriptionService) *PrescriptionHandler {
	return &PrescriptionHandler{
		prescriptionService: prescriptionService,
	}
}

func (h *PrescriptionHandler) CreatePrescription(c *fiber.Ctx) error {
	var req models.PrescriptionCreateRequest
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

	prescription, err := h.prescriptionService.CreatePrescription(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(prescription)
}

func (h *PrescriptionHandler) GetPrescriptions(c *fiber.Ctx) error {
	prescriptions, err := h.prescriptionService.GetAllPrescriptions()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch prescriptions",
		})
	}

	var response []models.PrescriptionResponse
	for _, p := range prescriptions {
		response = append(response, models.PrescriptionResponse{
			ID:          p.ID,
			PatientName: p.Patient.User.Name,
			DoctorName:  p.Doctor.Name,
			IssueDate:   p.IssueDate,
			ValidUntil:  p.ValidUntil,
			Status:      p.Status,
			Medications: p.Medications,
		})
	}

	return c.JSON(response)
}

func (h *PrescriptionHandler) GetPatientPrescriptions(c *fiber.Ctx) error {
	patientId, err := c.ParamsInt("patient_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid patient ID",
		})
	}

	prescriptions, err := h.prescriptionService.GetPatientPrescriptions(uint(patientId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch patient prescriptions",
		})
	}

	return c.JSON(prescriptions)
}

func (h *PrescriptionHandler) GetPrescription(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid prescription ID",
		})
	}

	prescription, err := h.prescriptionService.GetPrescriptionById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Prescription not found",
		})
	}

	return c.JSON(prescription)
}

func (h *PrescriptionHandler) DeletePrescription(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid prescription ID",
		})
	}

	if err := h.prescriptionService.DeletePrescription(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *PrescriptionHandler) UpdatePrescription(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid prescription ID",
		})
	}

	var req models.PrescriptionUpdateRequest
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

	prescription, err := h.prescriptionService.UpdatePrescription(uint(id), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(models.PrescriptionResponse{
		ID:          prescription.ID,
		PatientName: prescription.Patient.User.Name,
		DoctorName:  prescription.Doctor.Name,
		IssueDate:   prescription.IssueDate,
		ValidUntil:  prescription.ValidUntil,
		Status:      prescription.Status,
		Medications: prescription.Medications,
	})
}
