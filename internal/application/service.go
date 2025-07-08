package application

import "traefik-admin-go/internal/domain"

// Service represents the application service
type Service struct {
	domainService *domain.Service
}

// NewService creates a new application service
func NewService(domainService *domain.Service) *Service {
	return &Service{
		domainService: domainService,
	}
}

// GetHelloMessage returns a hello world message from the domain
func (s *Service) GetHelloMessage() string {
	return s.domainService.GetHelloMessage()
}