package vproductapi

import (
	"net/http"

	"github.com/tiberzus/goservice/api/sdk/http/mid"
	"github.com/tiberzus/goservice/app/domain/vproductapp"
	"github.com/tiberzus/goservice/app/sdk/auth"
	"github.com/tiberzus/goservice/app/sdk/authclient"
	"github.com/tiberzus/goservice/business/domain/userbus"
	"github.com/tiberzus/goservice/business/domain/vproductbus"
	"github.com/tiberzus/goservice/foundation/logger"
	"github.com/tiberzus/goservice/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log         *logger.Logger
	UserBus     *userbus.Business
	VProductBus *vproductbus.Business
	AuthClient  *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.Log, cfg.AuthClient)
	ruleAdmin := mid.Authorize(cfg.Log, cfg.AuthClient, auth.RuleAdminOnly)

	api := newAPI(vproductapp.NewApp(cfg.VProductBus))
	app.HandlerFunc(http.MethodGet, version, "/vproducts", api.query, authen, ruleAdmin)
}
