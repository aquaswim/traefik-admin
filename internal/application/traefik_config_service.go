package application

import (
	"traefik-admin-go/internal/domain"
)

type TraefikConfigService struct {
	routeRepository   domain.RouteRepository
	serviceRepository domain.ServiceRepository
}

func NewTraefikConfigService(
	routeRepository domain.RouteRepository,
	serviceRepository domain.ServiceRepository,
) *TraefikConfigService {
	return &TraefikConfigService{
		routeRepository:   routeRepository,
		serviceRepository: serviceRepository,
	}
}

func (s *TraefikConfigService) GetConfig() (*domain.TraefikConfig, error) {
	services, err := s.serviceRepository.GetAll()
	if err != nil {
		return nil, err
	}
	routes, err := s.routeRepository.GetAll()
	if err != nil {
		return nil, err
	}

	cfg, err := domain.NewTraefikConfig(services, routes)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
