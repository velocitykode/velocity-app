package routes

import (
	"velocity-app/internal/handlers"
	"velocity-app/internal/middleware"

	"github.com/velocitykode/velocity/pkg/router"
)

func init() {
	router.Register(func(r router.Router) {
		// Guest routes (only accessible when NOT authenticated)
		r.Group("", func(guest router.Router) {
			guest.Get("/login", handlers.AuthShowLoginForm)
			guest.Post("/login", handlers.AuthLogin)
			guest.Get("/register", handlers.AuthShowRegisterForm)
			guest.Post("/register", handlers.AuthRegister)
		}).Use(middleware.Guest)

		// Public routes
		r.Get("/health", handlers.Health)
		r.Post("/logout", handlers.AuthLogout)

		// Protected routes (require authentication)
		r.Group("", func(auth router.Router) {
			auth.Get("/", handlers.Dashboard)
			auth.Get("/dashboard", handlers.Dashboard)
		}).Use(middleware.Auth)
	})
}
