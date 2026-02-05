package middleware

import (
	"net/http"

	"github.com/velocitykode/velocity/pkg/auth"
	"github.com/velocitykode/velocity/pkg/router"
)

// Auth redirects to login if not authenticated
func Auth(next router.HandlerFunc) router.HandlerFunc {
	return func(ctx *router.Context) error {
		if !auth.Check(ctx.Request) {
			return ctx.Redirect(http.StatusSeeOther, "/login")
		}
		return next(ctx)
	}
}

// Guest redirects to dashboard if already authenticated
func Guest(next router.HandlerFunc) router.HandlerFunc {
	return func(ctx *router.Context) error {
		if auth.Check(ctx.Request) {
			return ctx.Redirect(http.StatusSeeOther, "/dashboard")
		}
		return next(ctx)
	}
}
