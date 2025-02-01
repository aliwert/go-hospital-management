package routes

import (
	"github.com/aliwert/go-hospital-management/internal/handlers"
	"github.com/aliwert/go-hospital-management/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, authHandler *handlers.AuthHandler) {
	// Public routes
	auth := app.Group("/api/v1/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// Protected routes
	api := app.Group("/api", middleware.AuthMiddleware())

	// Admin only routes
	admin := api.Group("/admin", middleware.RoleMiddleware("admin"))
	admin.Get("/users", authHandler.GetUsers)

	// Doctor routes
	doctor := api.Group("/doctor", middleware.RoleMiddleware("doctor"))
	doctor.Get("/profile", authHandler.GetProfile)

	// Patient routes
	patient := api.Group("/patient", middleware.RoleMiddleware("patient"))
	patient.Get("/profile", authHandler.GetProfile)
}
