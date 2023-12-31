// Package testgrp contains test handlers
package testgrp

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/sum28it/garage-service/foundation/web"
)

// Handlers manages the set of testgrp endpoints
type Handlers struct {
	DB *sqlx.DB
}

// Status represents a test handler for now.
func (h Handlers) Status(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// if n := rand.Intn(100); n%2 == 0 {
	// 	return v1Web.NewRequestError(errors.New("trusted error"), http.StatusBadRequest)
	// 	// 	return errors.New("ntrusted error")
	// 	// panic("Panicked")
	// }
	status := struct {
		Status string
	}{
		Status: "OK",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}
