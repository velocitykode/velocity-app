package middleware

import (
	"net/http"

	"github.com/velocitykode/velocity/pkg/router"
)

// ValidatePostSizeMiddleware rejects requests exceeding max body size
func ValidatePostSizeMiddleware(maxBytes int64) router.MiddlewareFunc {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *router.Context) error {
			// Only check for requests with a body
			if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
				c.Request.Body = http.MaxBytesReader(c.Response, c.Request.Body, maxBytes)
			}

			return next(c)
		}
	}
}
