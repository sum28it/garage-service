// Package handlers manages the different versions of the API.
package handlers

import (
	"context"
	"net/http"
	"os"

	"github.com/sum28it/garage-service/app/services/sales-api/handlers/v1/testgrp"
	"github.com/sum28it/garage-service/foundation/web"
	"go.uber.org/zap"
)

// A Handler is a type that handles a http request within our own little mini
// framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// // APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig) *web.App {
	app := web.NewApp(cfg.Shutdown)

	app.Handle(http.MethodGet, "/test", testgrp.Status)
	return app
}
