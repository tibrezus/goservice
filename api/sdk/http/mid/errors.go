package mid

import (
	"context"
	"net/http"

	"github.com/tiberzus/goservice/app/sdk/mid"
	"github.com/tiberzus/goservice/foundation/logger"
	"github.com/tiberzus/goservice/foundation/web"
)

// Errors executes the errors middleware functionality.
func Errors(log *logger.Logger) web.MidFunc {
	midFunc := func(ctx context.Context, r *http.Request, next mid.HandlerFunc) (mid.Encoder, error) {
		return mid.Errors(ctx, log, next)
	}

	return addMidFunc(midFunc)
}
