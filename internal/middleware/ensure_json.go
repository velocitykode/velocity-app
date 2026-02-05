package middleware

import (
	"github.com/velocitykode/velocity/pkg/router"
)

// EnsureJSONMiddleware ensures responses are JSON formatted for API routes
func EnsureJSONMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(c *router.Context) error {
		// Set JSON content type for all responses
		c.Response.Header().Set("Content-Type", "application/json")

		return next(c)
	}
}
