package repository

import (
	"errors"
	"sync"
	"traefik-admin-go/internal/domain"
)

// MemoryRouteRepository is an in-memory implementation of the RouteRepository interface
type MemoryRouteRepository struct {
	routes map[string]*domain.RouteModel
	mutex  sync.RWMutex
}

// NewMemoryRouteRepository creates a new domain.RouteRepository
func NewMemoryRouteRepository() domain.RouteRepository {
	return &MemoryRouteRepository{
		routes: make(map[string]*domain.RouteModel),
	}
}

// GetAll returns all routes
func (r *MemoryRouteRepository) GetAll() ([]*domain.RouteModel, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	routes := make([]*domain.RouteModel, 0, len(r.routes))
	for _, route := range r.routes {
		routes = append(routes, route)
	}
	return routes, nil
}

// GetByID returns a route by ID
func (r *MemoryRouteRepository) GetByID(id string) (*domain.RouteModel, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	route, exists := r.routes[id]
	if !exists {
		return nil, errors.New("route not found")
	}
	return route, nil
}

// Create creates a new route
func (r *MemoryRouteRepository) Create(route *domain.RouteModel) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.routes[route.ID]; exists {
		return errors.New("route already exists")
	}
	r.routes[route.ID] = route
	return nil
}

// Update updates an existing route
func (r *MemoryRouteRepository) Update(route *domain.RouteModel) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.routes[route.ID]; !exists {
		return errors.New("route not found")
	}
	r.routes[route.ID] = route
	return nil
}

// Delete deletes a route by ID
func (r *MemoryRouteRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.routes[id]; !exists {
		return errors.New("route not found")
	}
	delete(r.routes, id)
	return nil
}
