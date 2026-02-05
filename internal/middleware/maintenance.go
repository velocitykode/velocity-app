package middleware

import (
	"net/http"
	"os"

	"github.com/velocitykode/velocity/pkg/router"
)

// PreventRequestsDuringMaintenanceMiddleware returns 503 when app is in maintenance mode
func PreventRequestsDuringMaintenanceMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(c *router.Context) error {
		// Check if maintenance mode file exists
		if _, err := os.Stat("storage/framework/down"); err == nil {
			c.Response.WriteHeader(http.StatusServiceUnavailable)
			c.Response.Write([]byte("Service Temporarily Unavailable"))
			return nil
		}

		return next(c)
	}
}
