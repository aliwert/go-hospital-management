package routes

import (
	"github.com/aliwert/go-hospital-management/internal/database"
	"github.com/aliwert/go-hospital-management/internal/handlers"
	"github.com/aliwert/go-hospital-management/internal/middleware"
	"github.com/aliwert/go-hospital-management/internal/repositories"
	"github.com/aliwert/go-hospital-management/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router) {
	// Initialize auth components
	userRepo := repositories.NewUserRepository(database.GetDB())
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// Auth routes (public)
	auth := router.Group("/auth")
	{
		auth.Post("/register", authHandler.Register)
		auth.Post("/login", authHandler.Login)
	}

	// Protected routes
	api := router.Group("", middleware.AuthMiddleware())
	{
		api.Get("/profile", authHandler.GetProfile)

		// Admin routes
		admin := api.Group("/admin", middleware.RoleMiddleware("admin"))
		{
			admin.Get("/users", authHandler.GetUsers)
		}
	}
}
