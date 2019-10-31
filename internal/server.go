package internal

import (
	"fmt"
	"net/http"
	"strings"

	chilogger "github.com/766b/chi-logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/itross/sgul"
	"github.com/itross/sgulreg/internal/controllers"
	"github.com/itross/sgulreg/internal/services"
)

// Server defines the http server struct.
type Server struct {
	Registry services.Registry
	router   chi.Router
}

// NewServer returns a new Server instance.
func NewServer(r services.Registry) *Server {
	server := &Server{Registry: r}

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "OPTIONS", "DELETE", "HEAD"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	})

	server.router = chi.NewRouter()
	server.router.Use(
		cors.Handler,
		middleware.RequestID,
		middleware.RealIP,
		chilogger.NewZapMiddleware("router", sgul.GetLogger().Desugar()),
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	return server
}

// Serve confiures routes and starts the Http server.
func (s *Server) Serve() {
	defer func() {
		logger.Info("logger Sync")
		logger.Sync()
	}()

	// setup controllers
	registerController := controllers.NewRegisterController(s.Registry)
	logger.Info("controllers set up")

	apiConf := sgul.GetConfiguration().API
	s.router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("UP"))
	})

	// route countrollers
	s.router.Route(apiConf.Endpoint.BaseRoutingPath, func(r chi.Router) {
		r.Mount(registerController.Path, registerController.Router())
	})

	// log out configured routes
	walker := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		logger.Infow("initialized", "method", method, "route", route)
		return nil
	}

	if err := chi.Walk(s.router, walker); err != nil {
		logger.Panicf("error: %s\n", err.Error())
	}
	logger.Info("all api routes set up")

	// start the http server
	addr := fmt.Sprintf(":%d", apiConf.Endpoint.Port)
	logger.Infof("api http server up and running at http://localhost%s", addr)
	logger.Fatal(http.ListenAndServe(addr, s.router))
}
