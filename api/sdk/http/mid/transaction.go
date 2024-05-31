package mid

import (
	"context"
	"net/http"

	"github.com/tiberzus/goservice/app/sdk/mid"
	"github.com/tiberzus/goservice/business/sdk/transaction"
	"github.com/tiberzus/goservice/foundation/logger"
	"github.com/tiberzus/goservice/foundation/web"
)

// BeginCommitRollback executes the transaction middleware functionality.
func BeginCommitRollback(log *logger.Logger, bgn transaction.Beginner) web.MidFunc {
	midFunc := func(ctx context.Context, r *http.Request, next mid.HandlerFunc) (mid.Encoder, error) {
		return mid.BeginCommitRollback(ctx, log, bgn, next)
	}

	return addMidFunc(midFunc)
}
