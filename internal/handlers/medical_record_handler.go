package handlers

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"github.com/aliwert/go-hospital-management/internal/services"
	"github.com/aliwert/go-hospital-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type MedicalRecordHandler struct {
	medicalRecordService *services.MedicalRecordService
}

func NewMedicalRecordHandler(medicalRecordService *services.MedicalRecordService) *MedicalRecordHandler {
	return &MedicalRecordHandler{
		medicalRecordService: medicalRecordService,
	}
}

func (h *MedicalRecordHandler) CreateMedicalRecord(c *fiber.Ctx) error {
	var req models.MedicalRecordCreateRequest
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

	record, err := h.medicalRecordService.CreateMedicalRecord(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(record)
}

func (h *MedicalRecordHandler) GetMedicalRecords(c *fiber.Ctx) error {
	records, err := h.medicalRecordService.GetAllMedicalRecords()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch medical records",
		})
	}

	return c.JSON(records)
}

func (h *MedicalRecordHandler) GetMedicalRecord(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid medical record ID",
		})
	}

	record, err := h.medicalRecordService.GetMedicalRecordById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Medical record not found",
		})
	}

	return c.JSON(record)
}

func (h *MedicalRecordHandler) GetPatientMedicalRecords(c *fiber.Ctx) error {
	patientId, err := c.ParamsInt("patient_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid patient ID",
		})
	}

	records, err := h.medicalRecordService.GetPatientMedicalRecords(uint(patientId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch patient medical records",
		})
	}

	return c.JSON(records)
}

func (h *MedicalRecordHandler) UpdateMedicalRecord(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid medical record ID",
		})
	}

	record, err := h.medicalRecordService.GetMedicalRecordById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Medical record not found",
		})
	}

	if err := c.BodyParser(record); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.medicalRecordService.UpdateMedicalRecord(uint(id), record); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(record)
}

func (h *MedicalRecordHandler) DeleteMedicalRecord(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid medical record ID",
		})
	}

	if err := h.medicalRecordService.DeleteMedicalRecord(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
