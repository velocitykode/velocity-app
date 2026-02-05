package middleware

import (
	"os"

	"github.com/romsar/gonertia"
	"github.com/velocitykode/velocity/pkg/csrf"
	"github.com/velocitykode/velocity/pkg/router"
)

// CSRFTokenMiddleware sets the CSRF token in the template data for the meta tag.
// This allows Inertia/Axios to automatically read it and send with requests.
// Must run AFTER SessionMiddleware.
func CSRFTokenMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(c *router.Context) error {
		// Get session ID from cookie
		sessionName := os.Getenv("SESSION_NAME")
		if sessionName == "" {
			sessionName = "velocity_session"
		}

		ctx := c.Request.Context()

		sessionCookie, err := c.Request.Cookie(sessionName)
		if err == nil {
			// Get or generate CSRF token for this session
			token, err := csrf.GetGlobalToken(sessionCookie.Value)
			if err == nil && token != "" {
				// Set token in template data for the meta tag
				ctx = gonertia.SetTemplateDatum(ctx, "csrfToken", token)
			}
		}

		c.Request = c.Request.WithContext(ctx)
		return next(c)
	}
}
