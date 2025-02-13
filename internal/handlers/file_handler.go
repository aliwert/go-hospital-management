package handlers

import (
	"fmt"
	"io"

	"github.com/aliwert/go-hospital-management/internal/config"
	"github.com/aliwert/go-hospital-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type FileHandler struct{}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}

func (h *FileHandler) RegisterRoutes(app *fiber.App) {
	files := app.Group("/files")
	files.Post("/upload", h.UploadFile)
	files.Get("/download/:filename", h.DownloadFile)
	files.Delete("/delete/:filename", h.DeleteFile)
}

func (h *FileHandler) UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "File upload failed",
		})
	}

	fileContent, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not process file",
		})
	}
	defer fileContent.Close()

	err = utils.UploadToS3(config.GetS3Client(), fileContent, file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to upload to S3",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "File uploaded successfully",
	})
}

func (h *FileHandler) DownloadFile(c *fiber.Ctx) error {
	filename := c.Params("filename")

	output, err := utils.GetFromS3(config.GetS3Client(), filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to download file",
		})
	}
	defer output.Body.Close()

	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Set("Content-Type", "application/octet-stream")

	_, err = io.Copy(c.Response().BodyWriter(), output.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to stream file",
		})
	}

	return nil
}

func (h *FileHandler) DeleteFile(c *fiber.Ctx) error {
	filename := c.Params("filename")

	err := utils.DeleteFromS3(config.GetS3Client(), filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete file",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "File deleted successfully",
	})
}
