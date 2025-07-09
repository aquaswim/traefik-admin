package domain

import (
	"fmt"
	"strings"
)

type TraefikConfig struct {
	HTTP TraefikConfigData `json:"http,omitempty" yaml:"http,omitempty"`
	TCP  TraefikConfigData `json:"tcp,omitempty" yaml:"tcp,omitempty"`
	UDP  TraefikConfigData `json:"udp,omitempty" yaml:"udp,omitempty"`
}

type TraefikRouter struct {
	Rule    string `json:"rule,omitempty" yaml:"rule,omitempty"`
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}

type TraefikConfigData struct {
	Services map[string]TraefikService `json:"services,omitempty" yaml:"services,omitempty"`
	Routers  map[string]TraefikRouter  `json:"routers,omitempty" yaml:"routers,omitempty"`
}

type TraefikService struct {
	LoadBalancer *TraefikServiceLoadBalancer `json:"loadBalancer,omitempty" yaml:"loadBalancer,omitempty"`
	// other type TBA
}

type TraefikServiceLoadBalancer struct {
	Servers []TraefikServiceLoadBalancerServer `json:"servers,omitempty" yaml:"servers,omitempty"`
}

type TraefikServiceLoadBalancerServer struct {
	URL     string `json:"url,omitempty" yaml:"url,omitempty"`
	Address string `json:"address,omitempty" yaml:"address,omitempty"`
}

func NewTraefikConfig(
	services []*ServiceModel,
	routers []*RouteModel,
) (*TraefikConfig, error) {
	cfg := &TraefikConfig{
		HTTP: TraefikConfigData{
			Services: make(map[string]TraefikService),
			Routers:  make(map[string]TraefikRouter),
		},
		TCP: TraefikConfigData{
			Services: make(map[string]TraefikService),
			Routers:  make(map[string]TraefikRouter),
		},
		UDP: TraefikConfigData{
			Services: make(map[string]TraefikService),
			Routers:  make(map[string]TraefikRouter),
		},
	}
	for i := range services {
		switch strings.ToLower(services[i].Type) {
		case "http":
			cfg.HTTP.Services[services[i].ID] = TraefikService{
				LoadBalancer: &TraefikServiceLoadBalancer{
					Servers: []TraefikServiceLoadBalancerServer{
						{
							URL: services[i].Servers[0],
						},
					},
				},
			}
		case "tcp":
			cfg.TCP.Services[services[i].ID] = TraefikService{
				LoadBalancer: &TraefikServiceLoadBalancer{
					Servers: []TraefikServiceLoadBalancerServer{
						{
							Address: services[i].Servers[0],
						},
					},
				},
			}
		case "udp":
			cfg.UDP.Services[services[i].ID] = TraefikService{
				LoadBalancer: &TraefikServiceLoadBalancer{
					Servers: []TraefikServiceLoadBalancerServer{
						{
							Address: services[i].Servers[0],
						},
					},
				},
			}
		default:
			return nil, fmt.Errorf("service type %s is not supported", services[i].Type)
		}
	}
	for i := range routers {
		switch strings.ToLower(routers[i].Type) {
		case "http":
			cfg.HTTP.Routers[routers[i].ID] = TraefikRouter{
				Rule:    routers[i].Rule,
				Service: routers[i].Service,
			}
		case "tcp":
			cfg.TCP.Routers[routers[i].ID] = TraefikRouter{
				Rule:    routers[i].Rule,
				Service: routers[i].Service,
			}
		case "udp":
			cfg.UDP.Routers[routers[i].ID] = TraefikRouter{
				Rule:    routers[i].Rule,
				Service: routers[i].Service,
			}
		default:
			return nil, fmt.Errorf("route type %s is not supported", routers[i].Type)
		}
	}

	return cfg, nil
}
