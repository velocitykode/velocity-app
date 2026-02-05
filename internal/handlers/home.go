package handlers

import (
	"github.com/velocitykode/velocity/pkg/router"
	"github.com/velocitykode/velocity/pkg/view"
)

func Home(ctx *router.Context) error {
	view.Render(ctx.Response, ctx.Request, "Home", view.Props{
		"message": "Welcome to Velocity",
	})
	return nil
}
