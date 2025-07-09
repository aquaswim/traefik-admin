package domain

// RouteModel represents a route in the system
type RouteModel struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Rule    string `json:"rule"`
	Service string `json:"service"`
}

// NewRouteModel creates a new RouteModel
func NewRouteModel(id, routeType, rule, service string) *RouteModel {
	return &RouteModel{
		ID:      id,
		Type:    routeType,
		Rule:    rule,
		Service: service,
	}
}