package handlers

import (
	"github.com/velocitykode/velocity/router"
	"github.com/velocitykode/velocity/view"
)

func Home(ctx *router.Context) error {
	view.FromContext(ctx).Render(ctx.Response, ctx.Request, "Home", view.Props{
		"message": "Welcome to Velocity",
	})
	return nil
}
