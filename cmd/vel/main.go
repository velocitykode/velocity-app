// Package main provides the CLI entry point for Velocity projects.
// Build with: go build -o vel ./cmd/vel
// Production builds should use `go build .` which builds only the server.
package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/velocitykode/vel"

	// Import project packages to register with CLI
	_ "velocity-app/internal/app"
	_ "velocity-app/database/migrations"
)

func main() {
	// Load environment variables
	godotenv.Load()

	// Run CLI
	if err := vel.Execute(); err != nil {
		os.Exit(1)
	}
}
