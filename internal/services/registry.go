package services

import (
	"context"

	"github.com/ITResourcesOSS/sgulreg/internal/services/serializers"

	"github.com/go-chi/chi/middleware"

	"github.com/ITResourcesOSS/sgul"
	"github.com/ITResourcesOSS/sgulreg/internal/repositories"

	"github.com/ITResourcesOSS/sgulreg/internal/model"
	"github.com/ITResourcesOSS/sgulreg/pkg/sgulreg"
)

var logger = sgul.GetLogger().Sugar()

// Registry defines the interface to be implemented for a Service Registry.
type Registry interface {
	Register(ctx context.Context, r sgulreg.ServiceRegistrationRequest) (sgulreg.ServiceRegistrationResponse, error)
	Discovery(ctx context.Context, name string) (sgulreg.ServiceInfoResponse, error)
}

type registryService struct {
	serviceRepository repositories.ServiceRepository
}

// NewRegistry returns a new instance of a Registry Service.
func NewRegistry(sr repositories.ServiceRepository) Registry {
	return &registryService{serviceRepository: sr}
}

func (rs *registryService) Register(ctx context.Context, r sgulreg.ServiceRegistrationRequest) (sgulreg.ServiceRegistrationResponse, error) {
	requestID := middleware.GetReqID(ctx)
	service := model.NewService(r)

	logger.Infow("registering service instance", "instance", service.InstanceID, "request-id", requestID)
	if err := rs.serviceRepository.Save(ctx, service); err != nil {
		return sgulreg.ServiceRegistrationResponse{}, err
	}

	return serializers.NewServiceRegistrationResponse(service), nil
}

func (rs *registryService) Discovery(ctx context.Context, name string) (sgulreg.ServiceInfoResponse, error) {
	requestID := middleware.GetReqID(ctx)
	logger.Infow("discovering service", "service", name, "request-id", requestID)

	var instances []*model.Service
	var err error
	if instances, err = rs.serviceRepository.FindAllByServiceName(ctx, name); err != nil {
		return sgulreg.ServiceInfoResponse{}, err
	}

	return serializers.NewServiceInfoResponse(name, instances), nil
}
