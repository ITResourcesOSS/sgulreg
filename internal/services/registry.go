package services

import (
	"context"

	"github.com/itross/sgul/sgulreg"
	"github.com/itross/sgulreg/internal/services/serializers"

	"github.com/go-chi/chi/middleware"

	"github.com/itross/sgul"
	"github.com/itross/sgulreg/internal/repositories"

	"github.com/itross/sgulreg/internal/model"
)

var logger = sgul.GetLogger().Sugar()

// Registry defines the interface to be implemented for a Service Registry.
type Registry interface {
	Register(ctx context.Context, r sgulreg.ServiceRegistrationRequest) (sgulreg.ServiceRegistrationResponse, error)
	Discover(ctx context.Context, name string) (sgulreg.ServiceInfoResponse, error)
	DiscoverAll(ctx context.Context) ([]sgulreg.ServiceInfoResponse, error)
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

func (rs *registryService) Discover(ctx context.Context, name string) (sgulreg.ServiceInfoResponse, error) {
	requestID := middleware.GetReqID(ctx)
	logger.Infow("discovering service", "service", name, "request-id", requestID)

	var instances []*model.Service
	var err error
	if instances, err = rs.serviceRepository.FindAllByServiceName(ctx, name); err != nil {
		return sgulreg.ServiceInfoResponse{}, err
	}

	return serializers.NewServiceInfoResponse(name, instances), nil
}

func (rs *registryService) DiscoverAll(ctx context.Context) ([]sgulreg.ServiceInfoResponse, error) {
	requestID := middleware.GetReqID(ctx)
	logger.Infow("discovering all service", "request-id", requestID)

	// get all instances
	var instances []*model.Service
	var err error
	if instances, err = rs.serviceRepository.FindAll(ctx); err != nil {
		return []sgulreg.ServiceInfoResponse{}, err
	}

	// order all instances in a map by service-name
	tmpServices := make(map[string][]*model.Service)
	for _, i := range instances {
		if _, ok := tmpServices[i.Name]; !ok {
			tmpServices[i.Name] = []*model.Service{}
		}
		tmpServices[i.Name] = append(tmpServices[i.Name], i)
	}

	// serialize response
	response := make([]sgulreg.ServiceInfoResponse, len(tmpServices))
	idx := 0
	for k, v := range tmpServices {
		response[idx] = serializers.NewServiceInfoResponse(k, v)
		idx = idx + 1
		if idx > len(tmpServices)+1 {
			break
		}
	}

	return response, nil
}
