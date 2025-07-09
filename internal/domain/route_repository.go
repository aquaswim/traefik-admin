package domain

// RouteRepository defines the interface for route data storage
type RouteRepository interface {
	// GetAll returns all routes
	GetAll() ([]*RouteModel, error)
	
	// GetByID returns a route by ID
	GetByID(id string) (*RouteModel, error)
	
	// Create creates a new route
	Create(route *RouteModel) error
	
	// Update updates an existing route
	Update(route *RouteModel) error
	
	// Delete deletes a route by ID
	Delete(id string) error
}