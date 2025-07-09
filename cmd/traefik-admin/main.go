package main

import (
	"log"
	"traefik-admin-go/web"

	"github.com/gofiber/fiber/v2"
	"traefik-admin-go/internal/adapters/http"
	"traefik-admin-go/internal/adapters/repository"
	"traefik-admin-go/internal/application"
)

func main() {
	// Initialize repositories
	serviceRepository := repository.NewMemoryServiceRepository()
	routeRepository := repository.NewMemoryRouteRepository()

	// Initialize application layer
	serviceService := application.NewServiceService(serviceRepository)
	routeService := application.NewRouteService(routeRepository)
	traefikConfigService := application.NewTraefikConfigService(routeRepository, serviceRepository)

	// Initialize HTTP adapter (infrastructure layer)
	httpHandler := http.NewHandler(serviceService, routeService, traefikConfigService)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Traefik Admin",
	})

	// Register routes
	httpHandler.RegisterRoutes(app)
	web.RegisterRoutes(app)

	// Start server
	log.Println("Starting Traefik Admin server on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
