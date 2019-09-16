package internal

import (
	"log"

	"github.com/ITResourcesOSS/sgulreg/internal/registry"
)

// Application .
type Application struct {
	reg *registry.Registry
}

// NewApp .
func NewApp(reg *registry.Registry) *Application {
	return &Application{reg: reg}
}

// Start .
func (app *Application) Start() {
	log.Println("Application started...")
}
