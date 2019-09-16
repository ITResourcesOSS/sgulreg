package controllers

import (
	"net/http"

	"github.com/ITResourcesOSS/sgul"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var logger = sgul.GetLogger().Sugar()

// ServiceController defines the Service resource controller for the API.
type ServiceController struct {
	sgul.Controller
}

// NewServiceController returns a new ServiceController instance.
func NewServiceController() *ServiceController {
	return &ServiceController{Controller: sgul.NewController("/services")}
}

// Router returns routing paths for this ServiceController.
func (sc *ServiceController) Router() chi.Router {
	r := chi.NewRouter()
	r.Get("/", sc.all)
	return r
}

func (sc *ServiceController) all(w http.ResponseWriter, r *http.Request) {
	requestID := middleware.GetReqID(r.Context())
	logger.Infow("Request to get all Services information", "request-id", requestID)
}
