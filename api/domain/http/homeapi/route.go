package homeapi

import (
	"net/http"

	"github.com/tiberzus/goservice/api/sdk/http/mid"
	"github.com/tiberzus/goservice/app/domain/homeapp"
	"github.com/tiberzus/goservice/app/sdk/auth"
	"github.com/tiberzus/goservice/app/sdk/authclient"
	"github.com/tiberzus/goservice/business/domain/homebus"
	"github.com/tiberzus/goservice/business/domain/userbus"
	"github.com/tiberzus/goservice/foundation/logger"
	"github.com/tiberzus/goservice/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log        *logger.Logger
	UserBus    *userbus.Business
	HomeBus    *homebus.Business
	AuthClient *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.Log, cfg.AuthClient)
	ruleAny := mid.Authorize(cfg.Log, cfg.AuthClient, auth.RuleAny)
	ruleUserOnly := mid.Authorize(cfg.Log, cfg.AuthClient, auth.RuleUserOnly)
	ruleAuthorizeHome := mid.AuthorizeHome(cfg.Log, cfg.AuthClient, cfg.HomeBus)

	api := newAPI(homeapp.NewApp(cfg.HomeBus))
	app.HandlerFunc(http.MethodGet, version, "/homes", api.query, authen, ruleAny)
	app.HandlerFunc(http.MethodGet, version, "/homes/{home_id}", api.queryByID, authen, ruleAuthorizeHome)
	app.HandlerFunc(http.MethodPost, version, "/homes", api.create, authen, ruleUserOnly)
	app.HandlerFunc(http.MethodPut, version, "/homes/{home_id}", api.update, authen, ruleAuthorizeHome)
	app.HandlerFunc(http.MethodDelete, version, "/homes/{home_id}", api.delete, authen, ruleAuthorizeHome)
}
