package domain

// ServiceRepository defines the interface for service data storage
type ServiceRepository interface {
	// GetAll returns all services
	GetAll() ([]*ServiceModel, error)
	
	// GetByID returns a service by ID
	GetByID(id string) (*ServiceModel, error)
	
	// Create creates a new service
	Create(service *ServiceModel) error
	
	// Update updates an existing service
	Update(service *ServiceModel) error
	
	// Delete deletes a service by ID
	Delete(id string) error
}