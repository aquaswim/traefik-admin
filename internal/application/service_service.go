package application

import (
	"traefik-admin-go/internal/domain"
)

// ServiceService represents the application service for services
type ServiceService struct {
	serviceRepository domain.ServiceRepository
}

// NewServiceService creates a new service service
func NewServiceService(serviceRepository domain.ServiceRepository) *ServiceService {
	return &ServiceService{
		serviceRepository: serviceRepository,
	}
}

// GetAllServices returns all services
func (s *ServiceService) GetAllServices() ([]*domain.ServiceModel, error) {
	return s.serviceRepository.GetAll()
}

// GetServiceByID returns a service by ID
func (s *ServiceService) GetServiceByID(id string) (*domain.ServiceModel, error) {
	return s.serviceRepository.GetByID(id)
}

// CreateService creates a new service
func (s *ServiceService) CreateService(service *domain.ServiceModel) error {
	return s.serviceRepository.Create(service)
}

// UpdateService updates an existing service
func (s *ServiceService) UpdateService(service *domain.ServiceModel) error {
	return s.serviceRepository.Update(service)
}

// DeleteService deletes a service by ID
func (s *ServiceService) DeleteService(id string) error {
	return s.serviceRepository.Delete(id)
}