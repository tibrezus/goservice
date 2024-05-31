package tranapi

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/tiberzus/goservice/api/sdk/http/mid"
	"github.com/tiberzus/goservice/app/domain/tranapp"
	"github.com/tiberzus/goservice/app/sdk/auth"
	"github.com/tiberzus/goservice/app/sdk/authclient"
	"github.com/tiberzus/goservice/business/domain/productbus"
	"github.com/tiberzus/goservice/business/domain/userbus"
	"github.com/tiberzus/goservice/business/sdk/sqldb"
	"github.com/tiberzus/goservice/foundation/logger"
	"github.com/tiberzus/goservice/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log        *logger.Logger
	DB         *sqlx.DB
	UserBus    *userbus.Business
	ProductBus *productbus.Business
	AuthClient *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.Log, cfg.AuthClient)
	transaction := mid.BeginCommitRollback(cfg.Log, sqldb.NewBeginner(cfg.DB))
	ruleAdmin := mid.Authorize(cfg.Log, cfg.AuthClient, auth.RuleAdminOnly)

	api := newAPI(tranapp.NewApp(cfg.UserBus, cfg.ProductBus))
	app.HandlerFunc(http.MethodPost, version, "/tranexample", api.create, authen, ruleAdmin, transaction)
}
