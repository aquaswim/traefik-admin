package domain

// Service represents the domain service
type Service struct{}

// NewService creates a new domain service
func NewService() *Service {
	return &Service{}
}

// GetHelloMessage returns a hello world message
func (s *Service) GetHelloMessage() string {
	return "Hello, World!"
}