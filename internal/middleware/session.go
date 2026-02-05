package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/velocitykode/velocity/pkg/router"
)

// SessionMiddleware ensures a session cookie exists for every request.
// This must run BEFORE CSRF middleware so CSRF has a session ID to use.
func SessionMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(c *router.Context) error {
		cookieName := getSessionCookieName()

		// Check if session cookie already exists
		_, err := c.Request.Cookie(cookieName)
		if err == http.ErrNoCookie {
			// Create new session ID
			sessionID := generateSessionID()

			// Set cookie on response
			http.SetCookie(c.Response, &http.Cookie{
				Name:     cookieName,
				Value:    sessionID,
				Path:     getSessionPath(),
				Domain:   os.Getenv("SESSION_DOMAIN"),
				MaxAge:   getSessionLifetime() * 60, // Convert minutes to seconds
				HttpOnly: getSessionHttpOnly(),
				Secure:   getSessionSecure(),
				SameSite: getSessionSameSite(),
			})

			// Also add cookie to request so downstream handlers can read it
			c.Request.AddCookie(&http.Cookie{
				Name:  cookieName,
				Value: sessionID,
			})
		}

		return next(c)
	}
}

func generateSessionID() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return base64.URLEncoding.EncodeToString([]byte(time.Now().String()))
	}
	return base64.URLEncoding.EncodeToString(b)
}

func getSessionCookieName() string {
	name := os.Getenv("SESSION_NAME")
	if name == "" {
		return "velocity_session"
	}
	return name
}

func getSessionPath() string {
	path := os.Getenv("SESSION_PATH")
	if path == "" {
		return "/"
	}
	return path
}

func getSessionLifetime() int {
	lifetime := os.Getenv("SESSION_LIFETIME")
	if lifetime == "" {
		return 120 // 2 hours default
	}
	if val, err := strconv.Atoi(lifetime); err == nil {
		return val
	}
	return 120
}

func getSessionHttpOnly() bool {
	return os.Getenv("SESSION_HTTP_ONLY") != "false"
}

func getSessionSecure() bool {
	return os.Getenv("SESSION_SECURE") == "true"
}

func getSessionSameSite() http.SameSite {
	switch os.Getenv("SESSION_SAME_SITE") {
	case "strict":
		return http.SameSiteStrictMode
	case "none":
		return http.SameSiteNoneMode
	default:
		return http.SameSiteLaxMode
	}
}
