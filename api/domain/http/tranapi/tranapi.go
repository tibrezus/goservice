// Package tranapi maintains the web based api for tran access.
package tranapi

import (
	"context"
	"net/http"

	"github.com/tiberzus/goservice/app/domain/tranapp"
	"github.com/tiberzus/goservice/app/sdk/errs"
	"github.com/tiberzus/goservice/foundation/web"
)

type api struct {
	tranApp *tranapp.App
}

func newAPI(tranApp *tranapp.App) *api {
	return &api{
		tranApp: tranApp,
	}
}

func (api *api) create(ctx context.Context, r *http.Request) (web.Encoder, error) {
	var app tranapp.NewTran
	if err := web.Decode(r, &app); err != nil {
		return nil, errs.New(errs.FailedPrecondition, err)
	}

	prd, err := api.tranApp.Create(ctx, app)
	if err != nil {
		return nil, err
	}

	return prd, nil
}
