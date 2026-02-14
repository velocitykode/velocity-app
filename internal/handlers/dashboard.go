package handlers

import (
	"github.com/velocitykode/velocity/auth"
	"github.com/velocitykode/velocity/router"
	"github.com/velocitykode/velocity/view"
)

// Dashboard displays the dashboard
func Dashboard(ctx *router.Context) error {
	user := auth.FromContext(ctx).User(ctx.Request)

	// Convert user to map for props
	userMap := make(map[string]interface{})
	if authUser, ok := user.(*auth.AuthUser); ok {
		userMap["id"] = authUser.ID
		userMap["name"] = authUser.Name
		userMap["email"] = authUser.Email
	}

	view.FromContext(ctx).Render(ctx.Response, ctx.Request, "Dashboard", view.Props{
		"auth": map[string]interface{}{
			"user": userMap,
		},
	})
	return nil
}
