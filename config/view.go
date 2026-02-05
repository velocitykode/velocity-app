package config

import "github.com/velocitykode/velocity/pkg/config"

// GetViewTemplate returns the view template path (read at call time)
func GetViewTemplate() string {
	return config.Get("VIEW_TEMPLATE", "resources/views/app.go.html")
}

// GetViewVersion returns the view version (read at call time)
func GetViewVersion() string {
	return config.Get("VIEW_VERSION", "1.0")
}
