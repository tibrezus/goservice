package mid

import (
	"context"
	"net/http"

	"github.com/tiberzus/goservice/app/sdk/auth"
	"github.com/tiberzus/goservice/app/sdk/authclient"
	"github.com/tiberzus/goservice/app/sdk/mid"
	"github.com/tiberzus/goservice/business/domain/userbus"
	"github.com/tiberzus/goservice/foundation/logger"
	"github.com/tiberzus/goservice/foundation/web"
)

// Authenticate validates authentication via the auth service.
func Authenticate(log *logger.Logger, client *authclient.Client) web.MidFunc {
	midFunc := func(ctx context.Context, r *http.Request, next mid.HandlerFunc) (mid.Encoder, error) {
		return mid.Authenticate(ctx, log, client, r.Header.Get("authorization"), next)
	}

	return addMidFunc(midFunc)
}

// Bearer processes JWT authentication logic.
func Bearer(ath *auth.Auth) web.MidFunc {
	midFunc := func(ctx context.Context, r *http.Request, next mid.HandlerFunc) (mid.Encoder, error) {
		return mid.Bearer(ctx, ath, r.Header.Get("authorization"), next)
	}

	return addMidFunc(midFunc)
}

// Basic processes basic authentication logic.
func Basic(userBus *userbus.Business, ath *auth.Auth) web.MidFunc {
	midFunc := func(ctx context.Context, r *http.Request, next mid.HandlerFunc) (mid.Encoder, error) {
		return mid.Basic(ctx, ath, userBus, r.Header.Get("authorization"), next)
	}

	return addMidFunc(midFunc)
}
