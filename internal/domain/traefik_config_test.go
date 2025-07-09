package domain

import (
	"reflect"
	"testing"
)

func TestNewTraefikConfig(t *testing.T) {
	// Create test services for all types (HTTP, TCP, UDP)
	services := []*ServiceModel{
		NewServiceModel("http-service", "HTTP", []string{"http://server1:80", "http://server2:80"}),
		NewServiceModel("tcp-service", "TCP", []string{"tcp://server1:443", "tcp://server2:443"}),
		NewServiceModel("udp-service", "UDP", []string{"udp://server1:53", "udp://server2:53"}),
	}

	// Create test routes for all types (HTTP, TCP, UDP)
	routes := []*RouteModel{
		NewRouteModel("http-route", "HTTP", "Host(`example.com`)", "http-service"),
		NewRouteModel("tcp-route", "TCP", "HostSNI(`example.com`)", "tcp-service"),
		NewRouteModel("udp-route", "UDP", "HostSNI(`example.com`)", "udp-service"),
	}

	// Call the function under test
	cfg, err := NewTraefikConfig(services, routes)

	// Assertions
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify HTTP config
	if _, ok := cfg.HTTP.Services["http-service"]; !ok {
		t.Error("HTTP service not found in config")
	}
	expectedHTTPServer := []TraefikServiceLoadBalancerServer{
		{
			URL: "http://server1:80",
		},
	}
	if !reflect.DeepEqual(cfg.HTTP.Services["http-service"].LoadBalancer.Servers, expectedHTTPServer) {
		t.Error("HTTP service servers don't match")
	}
	if _, ok := cfg.HTTP.Routers["http-route"]; !ok {
		t.Error("HTTP router not found in config")
	}
	if cfg.HTTP.Routers["http-route"].Rule != "Host(`example.com`)" {
		t.Errorf("HTTP router rule incorrect, got %s", cfg.HTTP.Routers["http-route"].Rule)
	}
	if cfg.HTTP.Routers["http-route"].Service != "http-service" {
		t.Errorf("HTTP router service incorrect, got %s", cfg.HTTP.Routers["http-route"].Service)
	}

	// Verify TCP config
	if _, ok := cfg.TCP.Services["tcp-service"]; !ok {
		t.Error("TCP service not found in config")
	}
	expectedTCPServer := []TraefikServiceLoadBalancerServer{
		{
			Address: "tcp://server1:443",
		},
	}
	if !reflect.DeepEqual(cfg.TCP.Services["tcp-service"].LoadBalancer.Servers, expectedTCPServer) {
		t.Error("TCP service servers don't match")
	}
	if _, ok := cfg.TCP.Routers["tcp-route"]; !ok {
		t.Error("TCP router not found in config")
	}
	if cfg.TCP.Routers["tcp-route"].Rule != "HostSNI(`example.com`)" {
		t.Errorf("TCP router rule incorrect, got %s", cfg.TCP.Routers["tcp-route"].Rule)
	}
	if cfg.TCP.Routers["tcp-route"].Service != "tcp-service" {
		t.Errorf("TCP router service incorrect, got %s", cfg.TCP.Routers["tcp-route"].Service)
	}

	// Verify UDP config
	if _, ok := cfg.UDP.Services["udp-service"]; !ok {
		t.Error("UDP service not found in config")
	}
	expectedUDPServer := []TraefikServiceLoadBalancerServer{
		{
			Address: "udp://server1:53",
		},
	}
	if !reflect.DeepEqual(cfg.UDP.Services["udp-service"].LoadBalancer.Servers, expectedUDPServer) {
		t.Error("UDP service servers don't match")
	}
	if _, ok := cfg.UDP.Routers["udp-route"]; !ok {
		t.Error("UDP router not found in config")
	}
	if cfg.UDP.Routers["udp-route"].Rule != "HostSNI(`example.com`)" {
		t.Errorf("UDP router rule incorrect, got %s", cfg.UDP.Routers["udp-route"].Rule)
	}
	if cfg.UDP.Routers["udp-route"].Service != "udp-service" {
		t.Errorf("UDP router service incorrect, got %s", cfg.UDP.Routers["udp-route"].Service)
	}
}
