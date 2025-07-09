package application

import (
	"traefik-admin-go/internal/domain"
)

// RouteService represents the application service for routes
type RouteService struct {
	routeRepository domain.RouteRepository
}

// NewRouteService creates a new route service
func NewRouteService(routeRepository domain.RouteRepository) *RouteService {
	return &RouteService{
		routeRepository: routeRepository,
	}
}

// GetAllRoutes returns all routes
func (s *RouteService) GetAllRoutes() ([]*domain.RouteModel, error) {
	return s.routeRepository.GetAll()
}

// GetRouteByID returns a route by ID
func (s *RouteService) GetRouteByID(id string) (*domain.RouteModel, error) {
	return s.routeRepository.GetByID(id)
}

// CreateRoute creates a new route
func (s *RouteService) CreateRoute(route *domain.RouteModel) error {
	return s.routeRepository.Create(route)
}

// UpdateRoute updates an existing route
func (s *RouteService) UpdateRoute(route *domain.RouteModel) error {
	return s.routeRepository.Update(route)
}

// DeleteRoute deletes a route by ID
func (s *RouteService) DeleteRoute(id string) error {
	return s.routeRepository.Delete(id)
}