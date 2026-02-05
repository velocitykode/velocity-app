package app

import (
	"net/http"
	"os"

	"velocity-app/config"

	"github.com/joho/godotenv"
	"github.com/velocitykode/velocity/pkg/auth"
	"github.com/velocitykode/velocity/pkg/auth/drivers/guards"
	"github.com/velocitykode/velocity/pkg/crypto"
	"github.com/velocitykode/velocity/pkg/csrf"
	"github.com/velocitykode/velocity/pkg/csrf/stores"
	"github.com/velocitykode/velocity/pkg/log"
	"github.com/velocitykode/velocity/pkg/orm"
	"github.com/velocitykode/velocity/pkg/view"
)

func init() {
	godotenv.Load()
}

// Run starts the application
func Run() {
	log.Info("Velocity Inertia React Starter Kit")

	if err := initialize(); err != nil {
		log.Error("Failed to initialize application", "error", err)
		os.Exit(1)
	}

	httpKernel := New()
	httpKernel.Bootstrap()

	port := config.GetPort()
	log.Info("Server starting", "port", port)

	if err := http.ListenAndServe(":"+port, httpKernel.Handler()); err != nil {
		log.Error("Server failed to start", "error", err)
	}
}

// initialize bootstraps all application services
func initialize() error {
	initEvents()
	if err := initCrypto(); err != nil {
		return err
	}
	if err := orm.InitFromEnv(); err != nil {
		return err
	}
	if err := initAuth(); err != nil {
		return err
	}
	initCSRF()
	return initView()
}

func initCrypto() error {
	key := config.GetCryptoKey()
	if key != "" {
		return crypto.Init(crypto.Config{
			Key:    key,
			Cipher: config.GetCryptoCipher(),
		})
	}
	return nil
}

func initAuth() error {
	manager, err := auth.GetManager()
	if err != nil {
		return err
	}

	sessionConfig := auth.NewSessionConfigFromEnv()
	provider := auth.NewORMUserProvider(config.GetAuthModel())
	sessionGuard, err := guards.NewSessionGuard(provider, sessionConfig)
	if err != nil {
		return err
	}

	manager.RegisterGuard(config.GetAuthGuard(), sessionGuard)
	return nil
}

func initCSRF() {
	sessionName := os.Getenv("SESSION_NAME")
	if sessionName == "" {
		sessionName = "velocity_session"
	}

	csrfConfig := csrf.DefaultConfig()
	csrfConfig.Store = stores.NewSessionStore()
	csrfConfig.SessionCookieName = sessionName
	csrfConfig.ExcludePaths = []string{"/api/webhooks/*", "/health"}

	csrf.SetGlobalCSRF(csrf.New(csrfConfig))
}

func initView() error {
	template, err := view.LoadTemplateFromFile(config.GetViewTemplate())
	if err != nil {
		return err
	}

	if err := view.Initialize(view.Config{
		RootTemplate: template,
		Version:      config.GetViewVersion(),
	}); err != nil {
		return err
	}

	sessionName := os.Getenv("SESSION_NAME")
	if sessionName == "" {
		sessionName = "velocity_session"
	}

	view.SetSharePropsFunc(func(r *http.Request) (view.Props, error) {
		props := view.Props{}
		if cookie, err := r.Cookie(sessionName); err == nil {
			if token, err := csrf.GetGlobalToken(cookie.Value); err == nil && token != "" {
				props["csrf_token"] = token
			}
		}
		return props, nil
	})

	return nil
}
