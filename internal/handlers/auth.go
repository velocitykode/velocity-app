package handlers

import (
	"velocity-app/internal/models"

	"github.com/velocitykode/velocity/pkg/auth"
	"github.com/velocitykode/velocity/pkg/router"
	"github.com/velocitykode/velocity/pkg/view"
)

// AuthShowLoginForm displays the login page
func AuthShowLoginForm(ctx *router.Context) error {
	view.FromContext(ctx).Render(ctx.Response, ctx.Request, "Auth/Login", view.Props{})
	return nil
}

// AuthLogin handles the login request
func AuthLogin(ctx *router.Context) error {
	var formData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Remember bool   `json:"remember"`
	}

	if err := ctx.Bind(&formData); err != nil {
		// Fallback to form values
		formData.Email = ctx.Request.FormValue("email")
		formData.Password = ctx.Request.FormValue("password")
		formData.Remember = ctx.Request.FormValue("remember") == "on"
	}

	credentials := map[string]interface{}{
		"email":    formData.Email,
		"password": formData.Password,
	}

	authManager := auth.FromContext(ctx)
	success, _ := authManager.Attempt(ctx.Response, ctx.Request, credentials, formData.Remember)
	if !success {
		view.FromContext(ctx).Render(ctx.Response, ctx.Request, "Auth/Login", view.Props{
			"errors": map[string]string{
				"email": "These credentials do not match our records.",
			},
			"old": map[string]string{
				"email": formData.Email,
			},
		})
		return nil
	}

	view.FromContext(ctx).Redirect(ctx.Response, ctx.Request, "/dashboard")
	return nil
}

// AuthLogout handles the logout request
func AuthLogout(ctx *router.Context) error {
	auth.FromContext(ctx).Logout(ctx.Response, ctx.Request)
	view.FromContext(ctx).Redirect(ctx.Response, ctx.Request, "/login")
	return nil
}

// AuthShowRegisterForm displays the registration page
func AuthShowRegisterForm(ctx *router.Context) error {
	view.FromContext(ctx).Render(ctx.Response, ctx.Request, "Auth/Register", view.Props{})
	return nil
}

// AuthRegister handles the registration request
func AuthRegister(ctx *router.Context) error {
	var formData struct {
		Name                 string `json:"name"`
		Email                string `json:"email"`
		Password             string `json:"password"`
		PasswordConfirmation string `json:"password_confirmation"`
	}

	if err := ctx.Bind(&formData); err != nil {
		// Fallback to form values
		formData.Name = ctx.Request.FormValue("name")
		formData.Email = ctx.Request.FormValue("email")
		formData.Password = ctx.Request.FormValue("password")
		formData.PasswordConfirmation = ctx.Request.FormValue("password_confirmation")
	}

	// Validate required fields
	errors := make(map[string]string)
	if formData.Name == "" {
		errors["name"] = "Name is required."
	}
	if formData.Email == "" {
		errors["email"] = "Email is required."
	}
	if formData.Password == "" {
		errors["password"] = "Password is required."
	}

	viewEngine := view.FromContext(ctx)

	if len(errors) > 0 {
		viewEngine.Render(ctx.Response, ctx.Request, "Auth/Register", view.Props{
			"errors": errors,
			"old": map[string]string{
				"name":  formData.Name,
				"email": formData.Email,
			},
		})
		return nil
	}

	// Validate passwords match
	if formData.Password != formData.PasswordConfirmation {
		viewEngine.Render(ctx.Response, ctx.Request, "Auth/Register", view.Props{
			"errors": map[string]string{
				"password": "The password confirmation does not match.",
			},
			"old": map[string]string{
				"name":  formData.Name,
				"email": formData.Email,
			},
		})
		return nil
	}

	authManager := auth.FromContext(ctx)

	// Hash password
	hashedPassword, err := authManager.Hash(formData.Password)
	if err != nil {
		viewEngine.Render(ctx.Response, ctx.Request, "Auth/Register", view.Props{
			"errors": map[string]string{
				"password": "Failed to process password.",
			},
		})
		return nil
	}

	// Check if user already exists
	existingUser, _ := models.User{}.FindBy("email", formData.Email)
	if existingUser != nil {
		viewEngine.Render(ctx.Response, ctx.Request, "Auth/Register", view.Props{
			"errors": map[string]string{
				"email": "A user with this email already exists.",
			},
			"old": map[string]string{
				"name":  formData.Name,
				"email": formData.Email,
			},
		})
		return nil
	}

	// Create new user
	user, err := models.User{}.Create(map[string]any{
		"name":     formData.Name,
		"email":    formData.Email,
		"password": hashedPassword,
	})
	if err != nil {
		ctx.Log().Error("Failed to create user", "error", err)
		viewEngine.Render(ctx.Response, ctx.Request, "Auth/Register", view.Props{
			"errors": map[string]string{
				"email": "Failed to create account. Please try again.",
			},
		})
		return nil
	}

	ctx.Log().Info("User created successfully", "email", user.Email, "id", user.ID)

	// Auto-login the new user
	credentials := map[string]interface{}{
		"email":    formData.Email,
		"password": formData.Password,
	}

	success, _ := authManager.Attempt(ctx.Response, ctx.Request, credentials, false)
	if success {
		viewEngine.Redirect(ctx.Response, ctx.Request, "/dashboard")
	} else {
		viewEngine.Redirect(ctx.Response, ctx.Request, "/login")
	}
	return nil
}
