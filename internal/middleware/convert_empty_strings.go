package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"

	"github.com/velocitykode/velocity/pkg/router"
)

// ConvertEmptyStringsToNullMiddleware converts empty strings to nil in request data
func ConvertEmptyStringsToNullMiddleware(next router.HandlerFunc) router.HandlerFunc {
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

			// Convert empty strings to nil
			convertEmptyStrings(data)

			// Re-encode and set body
			convertedBody, _ := json.Marshal(data)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(convertedBody))
			c.Request.ContentLength = int64(len(convertedBody))
		}

		return next(c)
	}
}

// convertEmptyStrings recursively converts empty strings to nil in a map
func convertEmptyStrings(m map[string]interface{}) {
	for key, val := range m {
		switch v := val.(type) {
		case string:
			if v == "" {
				m[key] = nil
			}
		case map[string]interface{}:
			convertEmptyStrings(v)
		case []interface{}:
			for i, item := range v {
				if str, ok := item.(string); ok && str == "" {
					v[i] = nil
				} else if nested, ok := item.(map[string]interface{}); ok {
					convertEmptyStrings(nested)
				}
			}
		}
	}
}
