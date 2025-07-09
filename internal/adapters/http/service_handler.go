package http

import (
	"github.com/gofiber/fiber/v2"
	"traefik-admin-go/internal/application"
	"traefik-admin-go/internal/domain"
)

// ServiceHandler represents the HTTP handler for services
type ServiceHandler struct {
	serviceService *application.ServiceService
}

// NewServiceHandler creates a new service handler
func NewServiceHandler(serviceService *application.ServiceService) *ServiceHandler {
	return &ServiceHandler{
		serviceService: serviceService,
	}
}

// RegisterRoutes registers all service HTTP routes
func (h *ServiceHandler) RegisterRoutes(api fiber.Router) {
	services := api.Group("/services")
	services.Get("/", h.GetAllServices)
	services.Get("/:id", h.GetServiceByID)
	services.Post("/", h.CreateService)
	services.Put("/:id", h.UpdateService)
	services.Delete("/:id", h.DeleteService)
}

// GetAllServices handles the get all services request
func (h *ServiceHandler) GetAllServices(c *fiber.Ctx) error {
	services, err := h.serviceService.GetAllServices()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(services)
}

// GetServiceByID handles the get service by ID request
func (h *ServiceHandler) GetServiceByID(c *fiber.Ctx) error {
	id := c.Params("id")
	service, err := h.serviceService.GetServiceByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(service)
}

// CreateService handles the create service request
func (h *ServiceHandler) CreateService(c *fiber.Ctx) error {
	var service domain.ServiceModel
	if err := c.BodyParser(&service); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.serviceService.CreateService(&service); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(service)
}

// UpdateService handles the update service request
func (h *ServiceHandler) UpdateService(c *fiber.Ctx) error {
	id := c.Params("id")
	var service domain.ServiceModel
	if err := c.BodyParser(&service); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Ensure the ID in the URL matches the ID in the body
	service.ID = id

	if err := h.serviceService.UpdateService(&service); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(service)
}

// DeleteService handles the delete service request
func (h *ServiceHandler) DeleteService(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.serviceService.DeleteService(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
