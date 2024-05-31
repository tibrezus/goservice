package mid

import (
	"context"
	"net/http"

	"github.com/tiberzus/goservice/app/sdk/mid"
	"github.com/tiberzus/goservice/foundation/logger"
	"github.com/tiberzus/goservice/foundation/web"
)

// Logger executes the logger middleware functionality.
func Logger(log *logger.Logger) web.MidFunc {
	midFunc := func(ctx context.Context, r *http.Request, next mid.HandlerFunc) (mid.Encoder, error) {
		return mid.Logger(ctx, log, r.URL.Path, r.URL.RawQuery, r.Method, r.RemoteAddr, next)
	}

	return addMidFunc(midFunc)
}
