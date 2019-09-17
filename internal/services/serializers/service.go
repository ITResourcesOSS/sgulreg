package serializers

import (
	"time"

	"github.com/ITResourcesOSS/sgulreg/internal/controllers/dto"
	"github.com/ITResourcesOSS/sgulreg/internal/model"
)

// NewServiceRegistrationResponse returns a new registration response.
func NewServiceRegistrationResponse(s *model.Service) dto.ServiceRegistrationResponse {
	return dto.ServiceRegistrationResponse{
		InstanceID:            s.InstanceID,
		RegistrationTimestamp: time.Unix(s.RegistrationTimestamp, 0),
	}
}

// NewServiceInstanceInfo returns a new service instance info struct.
func NewServiceInstanceInfo(s *model.Service) dto.ServiceInstanceInfo {
	return dto.ServiceInstanceInfo{
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
func NewServiceInfoResponse(name string, instances []*model.Service) dto.ServiceInfoResponse {
	response := dto.ServiceInfoResponse{
		Name:      name,
		Instances: make([]dto.ServiceInstanceInfo, 0),
	}

	for _, service := range instances {
		response.Instances = append(response.Instances, NewServiceInstanceInfo(service))
	}

	return response
}
