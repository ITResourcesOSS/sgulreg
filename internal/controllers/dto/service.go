package dto

// ServiceRegistrationRequest defines the structure sent from service in order to be registered with th registry.
type ServiceRegistrationRequest struct {
	Name           string `json:"name"`
	Host           string `json:"host"`
	Schema         string `json:"schema"`
	InfoURL        string `json:"infoUrl"`
	HealthCheckURL string `json:"healthCheckUrl"`
}
