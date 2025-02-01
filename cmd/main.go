package main

import (
	"log"
	"os"

	"github.com/aliwert/go-hospital-management/internal/database"
	"github.com/aliwert/go-hospital-management/internal/handlers"
	"github.com/aliwert/go-hospital-management/internal/routes"
	"github.com/aliwert/go-hospital-management/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()

	app := fiber.New(fiber.Config{
		AppName: "Auth Service v1.0",
	})

	app.Use(logger.New())
	app.Use(cors.New())

	// Initialize services and handlers
	authService := services.NewAuthService(database.GetDB())
	authHandler := handlers.NewAuthHandler(authService)

	// Setup routes
	routes.SetupRoutes(app, authHandler)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Auth Service starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
