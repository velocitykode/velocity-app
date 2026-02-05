package app

import (
	"net/http"

	"github.com/velocitykode/velocity/pkg/router"
)

// HTTPKernel handles HTTP request lifecycle
type HTTPKernel struct {
	router *router.VelocityRouterV2
}

// New creates a new HTTP kernel
func New() *HTTPKernel {
	return &HTTPKernel{
		router: router.Get(),
	}
}

// Bootstrap sets up the HTTP layer
func (k *HTTPKernel) Bootstrap() {
	// Configure static file serving
	k.router.Static("public")

	// Get middleware stacks
	stacks := GetMiddlewareStacks()

	// Apply global middleware (runs for ALL requests)
	for _, middleware := range stacks.Global {
		k.router.Use(middleware)
	}

	// Apply web middleware (runs for browser/Inertia requests)
	// Note: These are applied globally for now
	// TODO: Apply only to web route group when route groups are implemented
	for _, middleware := range stacks.Web {
		k.router.Use(middleware)
	}

	// Load all registered routes
	router.LoadRoutes()
}

// Handler returns the HTTP handler
func (k *HTTPKernel) Handler() http.Handler {
	return k.router
}
