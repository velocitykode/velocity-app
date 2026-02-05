package middleware

import (
	"github.com/velocitykode/velocity/pkg/log"
	"github.com/velocitykode/velocity/pkg/router"
)

func LoggingMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(c *router.Context) error {
		log.Info("Request", "method", c.Request.Method, "path", c.Request.URL.Path)
		return next(c)
	}
}
