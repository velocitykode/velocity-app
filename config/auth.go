package config

import "github.com/velocitykode/velocity/pkg/config"

// GetAuthGuard returns the auth guard (read at call time)
func GetAuthGuard() string {
	return config.Get("AUTH_GUARD", "web")
}

// GetAuthModel returns the auth model (read at call time)
func GetAuthModel() string {
	return config.Get("AUTH_MODEL", "User")
}
