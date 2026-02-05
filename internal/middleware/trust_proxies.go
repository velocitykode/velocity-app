package middleware

import (
	"strings"

	"github.com/velocitykode/velocity/pkg/router"
)

// TrustProxiesMiddleware handles X-Forwarded-* headers from trusted proxies
func TrustProxiesMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(c *router.Context) error {
		// Handle X-Forwarded-For header
		if xff := c.Request.Header.Get("X-Forwarded-For"); xff != "" {
			// Get first IP from comma-separated list
			ips := strings.Split(xff, ",")
			if len(ips) > 0 {
				c.Request.RemoteAddr = strings.TrimSpace(ips[0])
			}
		}

		// Handle X-Forwarded-Proto header for scheme detection
		if proto := c.Request.Header.Get("X-Forwarded-Proto"); proto != "" {
			c.Request.URL.Scheme = proto
		}

		// Handle X-Forwarded-Host header
		if host := c.Request.Header.Get("X-Forwarded-Host"); host != "" {
			c.Request.Host = host
		}

		return next(c)
	}
}
