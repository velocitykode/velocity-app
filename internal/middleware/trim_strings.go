package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"

	"github.com/velocitykode/velocity/pkg/router"
)

// TrimStringsMiddleware trims whitespace from string inputs in form and JSON data
func TrimStringsMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(c *router.Context) error {
		// Only process POST, PUT, PATCH requests
		if c.Request.Method != "POST" && c.Request.Method != "PUT" && c.Request.Method != "PATCH" {
			return next(c)
		}

		// Handle JSON content
		if strings.Contains(c.Request.Header.Get("Content-Type"), "application/json") {
			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				return next(c)
			}
			defer c.Request.Body.Close()

			var data map[string]interface{}
			if err := json.Unmarshal(body, &data); err != nil {
				// Restore body if JSON parsing fails
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
				return next(c)
			}

			// Trim string values
			trimMapStrings(data)

			// Re-encode and set body
			trimmedBody, _ := json.Marshal(data)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(trimmedBody))
			c.Request.ContentLength = int64(len(trimmedBody))
		}

		return next(c)
	}
}

// trimMapStrings recursively trims strings in a map
func trimMapStrings(m map[string]interface{}) {
	for key, val := range m {
		switch v := val.(type) {
		case string:
			m[key] = strings.TrimSpace(v)
		case map[string]interface{}:
			trimMapStrings(v)
		case []interface{}:
			for i, item := range v {
				if str, ok := item.(string); ok {
					v[i] = strings.TrimSpace(str)
				} else if nested, ok := item.(map[string]interface{}); ok {
					trimMapStrings(nested)
				}
			}
		}
	}
}
