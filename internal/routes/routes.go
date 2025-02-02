package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "hospital-management",
		})
	})

	// API v1 routes
	v1 := app.Group("/api/v1")

	// Setup domain-specific routes
	SetupAuthRoutes(v1)
	SetupDoctorRoutes(v1)

	// Not found handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Endpoint not found",
		})
	})
}
