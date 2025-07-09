package domain

// ServiceModel represents a service in the system
type ServiceModel struct {
	ID      string   `json:"id"`
	Type    string   `json:"type"`
	Servers []string `json:"servers"`
}

// NewServiceModel creates a new ServiceModel
func NewServiceModel(id, serviceType string, servers []string) *ServiceModel {
	return &ServiceModel{
		ID:      id,
		Type:    serviceType,
		Servers: servers,
	}
}