package middleware

import (
	"net/http"

	"github.com/velocitykode/velocity/auth"
	"github.com/velocitykode/velocity/router"
)

// Auth redirects to login if not authenticated
func Auth(next router.HandlerFunc) router.HandlerFunc {
	return func(ctx *router.Context) error {
		if !auth.FromContext(ctx).Check(ctx.Request) {
			return ctx.Redirect(http.StatusSeeOther, "/login")
		}
		return next(ctx)
	}
}

// Guest redirects to dashboard if already authenticated
func Guest(next router.HandlerFunc) router.HandlerFunc {
	return func(ctx *router.Context) error {
		if auth.FromContext(ctx).Check(ctx.Request) {
			return ctx.Redirect(http.StatusSeeOther, "/dashboard")
		}
		return next(ctx)
	}
}
