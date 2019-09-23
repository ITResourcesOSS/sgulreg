package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ITResourcesOSS/sgul/sgulreg"
	"github.com/ITResourcesOSS/sgulreg/internal/services"

	"github.com/ITResourcesOSS/sgul"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

var logger = sgul.GetLogger().Sugar()

// RegisterController defines the Service resource controller for the API.
type RegisterController struct {
	sgul.Controller
	registry services.Registry
}

// NewRegisterController returns a new ServiceController instance.
func NewRegisterController(r services.Registry) *RegisterController {
	return &RegisterController{
		Controller: sgul.NewController("/services"),
		registry:   r,
	}
}

// Router returns routing paths for this ServiceController.
func (rc *RegisterController) Router() chi.Router {
	r := chi.NewRouter()
	r.Post("/", rc.register)
	r.Get("/", rc.all)
	r.Get("/{serviceName}", rc.discover)
	return r
}

func (rc *RegisterController) all(w http.ResponseWriter, r *http.Request) {
	requestID := middleware.GetReqID(r.Context())
	logger.Infow("Request to get all Services information", "request-id", requestID)

	var discoveryResponse []sgulreg.ServiceInfoResponse
	var err error
	if discoveryResponse, err = rc.registry.DiscoverAll(r.Context()); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, nil)
		return
	}
	if len(discoveryResponse) == 0 {
		render.Status(r, http.StatusNoContent)
		render.JSON(w, r, nil)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, discoveryResponse)
}

func (rc *RegisterController) register(w http.ResponseWriter, r *http.Request) {
	requestID := middleware.GetReqID(r.Context())
	logger.Infow("Request to register service instance for ", "request-id", requestID)

	registrationRequest := sgulreg.ServiceRegistrationRequest{}
	err := json.NewDecoder(r.Body).Decode(&registrationRequest)
	if err != nil {
		logger.Errorw("Error validating request", "error", err, "request-id", requestID)
		rc.RenderError(w, sgul.NewHTTPError(err, http.StatusBadRequest, "Unable to parse input request", requestID))
		return
	}

	logger.Infow("Service Registration Request",
		"name", registrationRequest.Name,
		"host", registrationRequest.Host,
		"schema", registrationRequest.Schema,
		"infoUrl", registrationRequest.InfoURL,
		"healthCheckUrl", registrationRequest.HealthCheckURL,
		"request-id", requestID)

	var response sgulreg.ServiceRegistrationResponse
	if response, err = rc.registry.Register(r.Context(), registrationRequest); err != nil {
		logger.Errorw("error registering service instance", "error", err, "request-id", requestID)
		rc.RenderError(w, sgul.NewHTTPError(err, http.StatusInternalServerError, "Unable to register service instance", requestID))
	} else {
		logger.Infow("service instance registered", "intanceId", response.InstanceID, "request-id", requestID)
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, response)
	}
}

func (rc *RegisterController) discover(w http.ResponseWriter, r *http.Request) {
	requestID := middleware.GetReqID(r.Context())
	serviceName := chi.URLParam(r, "serviceName")
	logger.Infow("Request to discover service", "serviceName", serviceName, "request-id", requestID)

	var discoveryResponse sgulreg.ServiceInfoResponse
	var err error
	if discoveryResponse, err = rc.registry.Discover(r.Context(), serviceName); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, nil)
		return
	}
	if len(discoveryResponse.Instances) == 0 {
		render.Status(r, http.StatusNoContent)
		render.JSON(w, r, nil)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, discoveryResponse)
}
