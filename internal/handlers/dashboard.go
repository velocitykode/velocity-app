package handlers

import (
	"github.com/velocitykode/velocity/pkg/auth"
	"github.com/velocitykode/velocity/pkg/router"
	"github.com/velocitykode/velocity/pkg/view"
)

// Dashboard displays the dashboard
func Dashboard(ctx *router.Context) error {
	user := auth.User(ctx.Request)

	// Convert user to map for props
	userMap := make(map[string]interface{})
	if authUser, ok := user.(*auth.AuthUser); ok {
		userMap["id"] = authUser.ID
		userMap["name"] = authUser.Name
		userMap["email"] = authUser.Email
	}

	view.Render(ctx.Response, ctx.Request, "Dashboard", view.Props{
		"auth": map[string]interface{}{
			"user": userMap,
		},
	})
	return nil
}
