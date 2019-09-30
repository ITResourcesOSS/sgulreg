package internal

import (
	"github.com/itross/sgul"
	"github.com/itross/sgulreg/internal/repositories"
	"github.com/itross/sgulreg/internal/services"
	"github.com/boltdb/bolt"
)

var logger = sgul.GetLogger().Sugar()

// Application .
type Application struct {
	db *bolt.DB
}

// NewApp .
func NewApp(db *bolt.DB) *Application {
	return &Application{db: db}
}

// Start .
func (app *Application) Start() {
	serviceRepository := repositories.NewServiceRepository(app.db)
	registry := services.NewRegistry(serviceRepository)
	server := NewServer(registry)
	logger.Info("http server set up. Start serving")
	server.Serve()
}
