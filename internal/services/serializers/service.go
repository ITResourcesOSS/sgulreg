package serializers

import (
	"time"

	"github.com/itross/sgul/registry"
	"github.com/itross/sgulreg/internal/model"
)

// NewServiceRegistrationResponse returns a new registration response.
func NewServiceRegistrationResponse(s *model.Service) registry.ServiceRegistrationResponse {
	return registry.ServiceRegistrationResponse{
		InstanceID:            s.InstanceID,
		RegistrationTimestamp: time.Unix(s.RegistrationTimestamp, 0),
	}
}

// NewServiceInstanceInfo returns a new service instance info struct.
func NewServiceInstanceInfo(s *model.Service) registry.ServiceInstanceInfo {
	return registry.ServiceInstanceInfo{
		InstanceID:            s.InstanceID,
		Host:                  s.Host,
		Schema:                s.Schema,
		InfoURL:               s.InfoURL,
		HealthCheckURL:        s.HealthCheckURL,
		RegistrationTimestamp: time.Unix(s.RegistrationTimestamp, 0),
		LastRefreshTimestamp:  time.Unix(s.LastRefreshTimestamp, 0),
	}
}

// NewServiceInfoResponse returns a new service discovery info response instance.
func NewServiceInfoResponse(name string, instances []*model.Service) registry.ServiceInfoResponse {
	response := registry.ServiceInfoResponse{
		Name:      name,
		Instances: make([]registry.ServiceInstanceInfo, 0),
	}

	for _, service := range instances {
		response.Instances = append(response.Instances, NewServiceInstanceInfo(service))
	}

	return response
}
