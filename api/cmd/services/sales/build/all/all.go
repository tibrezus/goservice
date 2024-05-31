// Package all binds all the routes into the specified app.
package all

import (
	"time"

	"github.com/tiberzus/goservice/api/domain/http/checkapi"
	"github.com/tiberzus/goservice/api/domain/http/homeapi"
	"github.com/tiberzus/goservice/api/domain/http/productapi"
	"github.com/tiberzus/goservice/api/domain/http/rawapi"
	"github.com/tiberzus/goservice/api/domain/http/tranapi"
	"github.com/tiberzus/goservice/api/domain/http/userapi"
	"github.com/tiberzus/goservice/api/domain/http/vproductapi"
	"github.com/tiberzus/goservice/api/sdk/http/mux"
	"github.com/tiberzus/goservice/business/domain/homebus"
	"github.com/tiberzus/goservice/business/domain/homebus/stores/homedb"
	"github.com/tiberzus/goservice/business/domain/productbus"
	"github.com/tiberzus/goservice/business/domain/productbus/stores/productdb"
	"github.com/tiberzus/goservice/business/domain/userbus"
	"github.com/tiberzus/goservice/business/domain/userbus/stores/usercache"
	"github.com/tiberzus/goservice/business/domain/userbus/stores/userdb"
	"github.com/tiberzus/goservice/business/domain/vproductbus"
	"github.com/tiberzus/goservice/business/domain/vproductbus/stores/vproductdb"
	"github.com/tiberzus/goservice/business/sdk/delegate"
	"github.com/tiberzus/goservice/foundation/web"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg mux.Config) {

	// Construct the business domain packages we need here so we are using the
	// sames instances for the different set of domain apis.
	delegate := delegate.New(cfg.Log)
	userBus := userbus.NewBusiness(cfg.Log, delegate, usercache.NewStore(cfg.Log, userdb.NewStore(cfg.Log, cfg.DB), time.Hour))
	productBus := productbus.NewBusiness(cfg.Log, userBus, delegate, productdb.NewStore(cfg.Log, cfg.DB))
	homeBus := homebus.NewBusiness(cfg.Log, userBus, delegate, homedb.NewStore(cfg.Log, cfg.DB))
	vproductBus := vproductbus.NewBusiness(vproductdb.NewStore(cfg.Log, cfg.DB))

	checkapi.Routes(app, checkapi.Config{
		Build: cfg.Build,
		Log:   cfg.Log,
		DB:    cfg.DB,
	})

	homeapi.Routes(app, homeapi.Config{
		Log:        cfg.Log,
		UserBus:    userBus,
		HomeBus:    homeBus,
		AuthClient: cfg.AuthClient,
	})

	productapi.Routes(app, productapi.Config{
		Log:        cfg.Log,
		UserBus:    userBus,
		ProductBus: productBus,
		AuthClient: cfg.AuthClient,
	})

	rawapi.Routes(app)

	tranapi.Routes(app, tranapi.Config{
		Log:        cfg.Log,
		DB:         cfg.DB,
		UserBus:    userBus,
		ProductBus: productBus,
		AuthClient: cfg.AuthClient,
	})

	userapi.Routes(app, userapi.Config{
		Log:        cfg.Log,
		UserBus:    userBus,
		AuthClient: cfg.AuthClient,
	})

	vproductapi.Routes(app, vproductapi.Config{
		Log:         cfg.Log,
		UserBus:     userBus,
		VProductBus: vproductBus,
		AuthClient:  cfg.AuthClient,
	})
}
