package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"traefik-admin-go/internal/adapters/http"
	"traefik-admin-go/internal/application"
	"traefik-admin-go/internal/domain"
)

func main() {
	// Initialize domain layer
	domainService := domain.NewService()

	// Initialize application layer
	appService := application.NewService(domainService)

	// Initialize HTTP adapter (infrastructure layer)
	httpHandler := http.NewHandler(appService)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Traefik Admin",
	})

	// Register routes
	httpHandler.RegisterRoutes(app)

	// Start server
	log.Println("Starting Traefik Admin server on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}