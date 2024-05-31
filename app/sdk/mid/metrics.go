package mid

import (
	"context"

	"github.com/tiberzus/goservice/app/sdk/metrics"
)

// Metrics updates program counters.
func Metrics(ctx context.Context, next HandlerFunc) (Encoder, error) {
	ctx = metrics.Set(ctx)

	resp, err := next(ctx)

	n := metrics.AddRequests(ctx)

	if n%1000 == 0 {
		metrics.AddGoroutines(ctx)
	}

	if err != nil {
		metrics.AddErrors(ctx)
	}

	return resp, err
}
