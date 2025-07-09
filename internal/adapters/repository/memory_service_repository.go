package repository

import (
	"errors"
	"sync"
	"traefik-admin-go/internal/domain"
)

// MemoryServiceRepository is an in-memory implementation of the ServiceRepository interface
type MemoryServiceRepository struct {
	services map[string]*domain.ServiceModel
	mutex    sync.RWMutex
}

// NewMemoryServiceRepository creates a new domain.ServiceRepository
func NewMemoryServiceRepository() domain.ServiceRepository {
	return &MemoryServiceRepository{
		services: make(map[string]*domain.ServiceModel),
	}
}

// GetAll returns all services
func (r *MemoryServiceRepository) GetAll() ([]*domain.ServiceModel, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	services := make([]*domain.ServiceModel, 0, len(r.services))
	for _, service := range r.services {
		services = append(services, service)
	}
	return services, nil
}

// GetByID returns a service by ID
func (r *MemoryServiceRepository) GetByID(id string) (*domain.ServiceModel, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	service, exists := r.services[id]
	if !exists {
		return nil, errors.New("service not found")
	}
	return service, nil
}

// Create creates a new service
func (r *MemoryServiceRepository) Create(service *domain.ServiceModel) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.services[service.ID]; exists {
		return errors.New("service already exists")
	}
	r.services[service.ID] = service
	return nil
}

// Update updates an existing service
func (r *MemoryServiceRepository) Update(service *domain.ServiceModel) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.services[service.ID]; !exists {
		return errors.New("service not found")
	}
	r.services[service.ID] = service
	return nil
}

// Delete deletes a service by ID
func (r *MemoryServiceRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.services[id]; !exists {
		return errors.New("service not found")
	}
	delete(r.services, id)
	return nil
}
