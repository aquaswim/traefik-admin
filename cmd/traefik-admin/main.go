package main

import (
	"fmt"
	"github.com/dgraph-io/badger/v4"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"traefik-admin-go/web"

	"github.com/gofiber/fiber/v2"
	"traefik-admin-go/internal/adapters/http"
	"traefik-admin-go/internal/adapters/repository"
	"traefik-admin-go/internal/application"
)

func main() {
	// db
	db, err := badger.Open(badger.DefaultOptions("./dev.db"))
	if err != nil {
		log.Fatal(err)
	}

	// Initialize repositories with Badger
	serviceRepository := repository.NewBadgerServiceRepository(db)
	routeRepository := repository.NewBadgerRouteRepository(db)

	// Initialize application layer
	serviceService := application.NewServiceService(serviceRepository)
	routeService := application.NewRouteService(routeRepository)
	traefikConfigService := application.NewTraefikConfigService(routeRepository, serviceRepository)

	// Initialize HTTP adapter (infrastructure layer)
	httpHandler := http.NewHandler(serviceService, routeService, traefikConfigService)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:     "Traefik Admin",
		IdleTimeout: 5 * time.Second,
	})

	// Register routes
	httpHandler.RegisterRoutes(app)
	web.RegisterRoutes(app)

	// Start server
	log.Println("Starting Traefik Admin server on :3000")
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	err = db.Close()
	if err != nil {
		log.Printf("Error closing db: %v", err)
	}
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
}
