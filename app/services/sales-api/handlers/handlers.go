// Package handlers manages the different versions of the API.
package handlers

import (
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/sum28it/garage-service/app/services/sales-api/handlers/v1/testgrp"
	"github.com/sum28it/garage-service/business/web/auth"
	"github.com/sum28it/garage-service/business/web/v1/mid"
	"github.com/sum28it/garage-service/foundation/web"
	"go.uber.org/zap"
)

// // APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
	Auth     *auth.Auth
	DB       *sqlx.DB
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig) http.Handler {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Metrics(), mid.Panics())

	tg := testgrp.Handlers{
		DB: cfg.DB,
	}
	app.Handle(http.MethodGet, "/status", tg.Status)
	app.Handle(http.MethodGet, "/auth", tg.Status, mid.Authenticate(cfg.Auth), mid.Authorize(cfg.Auth, auth.RuleAdminOnly))
	return app
}
