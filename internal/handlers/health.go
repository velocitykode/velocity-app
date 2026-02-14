package handlers

import (
	"github.com/velocitykode/velocity/router"
)

func Health(ctx *router.Context) error {
	return ctx.JSON(router.StatusOK, map[string]string{
		"status": "healthy",
	})
}
