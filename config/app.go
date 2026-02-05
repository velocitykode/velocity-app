package config

import "github.com/velocitykode/velocity/pkg/config"

// GetAppName returns the app name (read at call time)
func GetAppName() string {
	return config.Get("APP_NAME", "Velocity")
}

// GetAppEnv returns the app environment (read at call time)
func GetAppEnv() string {
	return config.Get("APP_ENV", "development")
}

// GetPort returns the port (read at call time)
func GetPort() string {
	return config.Get("PORT", "4000")
}
