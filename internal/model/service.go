package model

import (
	"fmt"
	"time"

	"github.com/ITResourcesOSS/sgul/sgulreg"
)

// Service is the service info struct to be saved in registry.
type Service struct {
	InstanceID            string
	Name                  string
	Host                  string
	Schema                string
	InfoURL               string
	HealthCheckURL        string
	RegistrationTimestamp int64
	LastRefreshTimestamp  int64
}

// NewService returns a new instance of the Service model from Service registration request.
func NewService(r sgulreg.ServiceRegistrationRequest) *Service {
	timestamp := time.Now().Unix()
	return &Service{
		InstanceID:            fmt.Sprintf("%s@%s", r.Name, r.Host),
		Name:                  r.Name,
		Host:                  r.Host,
		Schema:                r.Schema,
		InfoURL:               r.InfoURL,
		HealthCheckURL:        r.HealthCheckURL,
		RegistrationTimestamp: timestamp,
		LastRefreshTimestamp:  timestamp,
	}
}
