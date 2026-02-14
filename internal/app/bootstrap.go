package app

import (
	"net/http"
	"os"

	"velocity-app/config"

	"github.com/velocitykode/velocity"
	"github.com/velocitykode/velocity/auth"
	"github.com/velocitykode/velocity/auth/drivers/guards"
	"github.com/velocitykode/velocity/csrf"
	"github.com/velocitykode/velocity/csrf/stores"
	"github.com/velocitykode/velocity/view"
)

// Bootstrap configures app-specific services on the Velocity app instance.
// Core services (crypto, ORM, logger, cache, events) are already initialized
// by velocity.New().
func Bootstrap(v *velocity.App) error {
	// 1. Register auth guards (app-specific: session guard with user model)
	if err := bootstrapAuth(v); err != nil {
		return err
	}

	// 2. Configure CSRF with session store (app-specific)
	bootstrapCSRF(v)

	// 3. Configure view engine with template and shared props (app-specific)
	if err := bootstrapView(v); err != nil {
		return err
	}

	// 4. Apply middleware to the router
	bootstrapMiddleware(v)

	// 5. Serve static files
	v.Router.Static("public")

	// 6. Register event listeners (app-specific)
	initEvents(v.Log, v.Events)

	return nil
}

func bootstrapAuth(v *velocity.App) error {
	sessionConfig := auth.NewSessionConfigFromEnv()
	provider := auth.NewORMUserProvider(v.DB.DB(), config.GetAuthModel(), nil)
	sessionGuard, err := guards.NewSessionGuard(provider, sessionConfig, v.Crypto)
	if err != nil {
		return err
	}

	v.Auth.(*auth.Manager).RegisterGuard(config.GetAuthGuard(), sessionGuard)
	return nil
}

func bootstrapCSRF(v *velocity.App) {
	sessionName := os.Getenv("SESSION_NAME")
	if sessionName == "" {
		sessionName = "velocity_session"
	}

	csrfConfig := csrf.DefaultConfig()
	csrfConfig.Store = stores.NewSessionStore()
	csrfConfig.SessionCookieName = sessionName
	csrfConfig.ExcludePaths = []string{"/api/webhooks/*", "/health"}

	// Replace the default CSRF instance with our configured one
	v.CSRF = csrf.New(csrfConfig)
}

func bootstrapView(v *velocity.App) error {
	template, err := view.LoadTemplateFromFile(config.GetViewTemplate())
	if err != nil {
		return err
	}

	engine, err := view.NewEngine(view.Config{
		RootTemplate: template,
		Version:      config.GetViewVersion(),
	})
	if err != nil {
		return err
	}

	// Set on the app instance
	v.View = engine

	sessionName := os.Getenv("SESSION_NAME")
	if sessionName == "" {
		sessionName = "velocity_session"
	}

	csrfInstance := v.CSRF.(*csrf.CSRF)
	engine.SetSharePropsFunc(func(r *http.Request) (view.Props, error) {
		props := view.Props{}
		if cookie, err := r.Cookie(sessionName); err == nil {
			if token, err := csrfInstance.GetToken(cookie.Value); err == nil && token != "" {
				props["csrf_token"] = token
			}
		}
		return props, nil
	})

	return nil
}

func bootstrapMiddleware(v *velocity.App) {
	stacks := GetMiddlewareStacks(v)

	for _, mw := range stacks.Global {
		v.Router.Use(mw)
	}
	for _, mw := range stacks.Web {
		v.Router.Use(mw)
	}
}
