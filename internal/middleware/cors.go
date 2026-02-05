package middleware

import (
	"net/http"

	"github.com/velocitykode/velocity/pkg/router"
)

// CORSMiddleware handles CORS preflight and headers
func CORSMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(c *router.Context) error {
		c.Response.Header().Set("Access-Control-Allow-Origin", "*")
		c.Response.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, X-Inertia, X-Inertia-Version")
		c.Response.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.Response.WriteHeader(http.StatusOK)
			return nil
		}

		return next(c)
	}
}
