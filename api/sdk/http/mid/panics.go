package mid

import (
	"context"
	"net/http"

	"github.com/tiberzus/goservice/app/sdk/mid"
	"github.com/tiberzus/goservice/foundation/web"
)

// Panics executes the panic middleware functionality.
func Panics() web.MidFunc {
	midFunc := func(ctx context.Context, r *http.Request, next mid.HandlerFunc) (mid.Encoder, error) {
		return mid.Panics(ctx, next)
	}

	return addMidFunc(midFunc)
}
