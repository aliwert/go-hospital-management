package routes

import (
	"github.com/aliwert/go-hospital-management/internal/handlers"
	"github.com/aliwert/go-hospital-management/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, authHandler *handlers.AuthHandler) {
	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "auth-service",
		})
	})

	// API v1 routes
	v1 := app.Group("/api/v1")

	// Auth routes (public)
	auth := v1.Group("/auth")
	{
		auth.Post("/register", authHandler.Register)
		auth.Post("/login", authHandler.Login)
	}

	// Protected routes
	api := v1.Group("", middleware.AuthMiddleware())
	{
		// User profile route (accessible by all authenticated users)
		api.Get("/profile", authHandler.GetProfile)

		// Admin routes
		admin := api.Group("/admin", middleware.RoleMiddleware("admin"))
		{
			admin.Get("/users", authHandler.GetUsers)
		}
	}

	// Not found handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Endpoint not found",
		})
	})
}
