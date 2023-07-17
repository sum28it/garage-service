package mid

import (
	"context"
	"net/http"

	"github.com/sum28it/garage-service/foundation/web"
	"go.uber.org/zap"
)

// Logger retuns the logging middleware that logs information about the request
func Logger(log *zap.SugaredLogger) web.Middleware {

	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			log.Infow("request started", "method", r.Method, "path", r.URL.Path, "remote address", r.RemoteAddr)

			err := handler(ctx, w, r)

			log.Infow("request completed", "method", r.Method, "path", r.URL.Path, "remote address", r.RemoteAddr)

			return err
		}

		return h
	}

	return m

}
