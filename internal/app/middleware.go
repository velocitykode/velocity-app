package app

import (
	"velocity-app/internal/middleware"

	"github.com/velocitykode/velocity/pkg/csrf"
	"github.com/velocitykode/velocity/pkg/router"
	"github.com/velocitykode/velocity/pkg/view"
)

// MiddlewareStacks defines all middleware stacks for the application
type MiddlewareStacks struct {
	Global []router.MiddlewareFunc
	Web    []router.MiddlewareFunc
	API    []router.MiddlewareFunc
}

// GetMiddlewareStacks returns configured middleware stacks
func GetMiddlewareStacks() *MiddlewareStacks {
	return &MiddlewareStacks{
		Global: globalMiddleware(),
		Web:    webMiddleware(),
		API:    apiMiddleware(),
	}
}

// globalMiddleware returns middleware that runs for ALL requests
// These run before any route-specific or group-specific middleware
func globalMiddleware() []router.MiddlewareFunc {
	return []router.MiddlewareFunc{
		middleware.LoggingMiddleware,                      // Log all requests
		middleware.TrustProxiesMiddleware,                 // Handle X-Forwarded-* headers
		middleware.CORSMiddleware,                         // Handle CORS preflight and headers
		middleware.PreventRequestsDuringMaintenanceMiddleware, // Return 503 when in maintenance mode
		middleware.ValidatePostSizeMiddleware(10 << 20),   // Reject requests > 10MB
		middleware.TrimStringsMiddleware,                  // Trim whitespace from string inputs
		middleware.ConvertEmptyStringsToNullMiddleware,    // Convert "" to nil for cleaner handling
	}
}

// webMiddleware returns middleware for browser/web requests (Inertia, HTML)
// These run only for routes in the "web" group
func webMiddleware() []router.MiddlewareFunc {
	return []router.MiddlewareFunc{
		middleware.SessionMiddleware,    // Create session cookie (must be before CSRF)
		middleware.CSRFTokenMiddleware,  // Set CSRF token in template data
		csrf.Middleware(), // CSRF protection (validates token)
		view.Middleware(),               // Inertia middleware - handles X-Inertia headers, version, etc.
	}
}

// apiMiddleware returns middleware for API requests (JSON APIs)
// These run only for routes in the "api" group
func apiMiddleware() []router.MiddlewareFunc {
	return []router.MiddlewareFunc{
		middleware.EnsureJSONMiddleware, // Ensure response is JSON formatted
		// Throttle - Rate limiting (requires framework support)
		// SubstituteBindings - Route model bindings (requires framework support)
	}
}
