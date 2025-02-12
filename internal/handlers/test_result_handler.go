package handlers

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"github.com/aliwert/go-hospital-management/internal/services"
	"github.com/gofiber/fiber/v2"
)

type TestResultHandler struct {
	testResultService *services.TestResultService
}

func NewTestResultHandler(testResultService *services.TestResultService) *TestResultHandler {
	return &TestResultHandler{
		testResultService: testResultService,
	}
}

func (h *TestResultHandler) CreateTestResult(c *fiber.Ctx) error {
	var result models.TestResult
	if err := c.BodyParser(&result); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.testResultService.CreateTestResult(&result); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func (h *TestResultHandler) GetTestResult(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid test result ID",
		})
	}

	result, err := h.testResultService.GetTestResultById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Test result not found",
		})
	}

	return c.JSON(result)
}

func (h *TestResultHandler) GetTestResultsByMedicalRecord(c *fiber.Ctx) error {
	medicalRecordId, err := c.ParamsInt("medical_record_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid medical record ID",
		})
	}

	results, err := h.testResultService.GetTestResultsByMedicalRecordId(uint(medicalRecordId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch test results",
		})
	}

	return c.JSON(results)
}

func (h *TestResultHandler) UpdateTestResult(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid test result ID",
		})
	}

	var result models.TestResult
	if err := c.BodyParser(&result); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result.ID = uint(id)
	if err := h.testResultService.UpdateTestResult(&result); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(result)
}

func (h *TestResultHandler) DeleteTestResult(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid test result ID",
		})
	}

	if err := h.testResultService.DeleteTestResult(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
