package all

import (
	authapi "github.com/bentenison/microservice/api/domain/auth-api"
	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/domain/authbus"
	"github.com/bentenison/microservice/business/domain/authbus/stores/authdb"
	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/foundation/web"
)

func Routes() add {
	return add{}
}

type add struct{}

func (a add) Add(app *web.App, cfg mux.Config) {
	delegate := delegate.New(cfg.Log)
	authbus := authbus.NewBusiness(cfg.Log, delegate, cfg.DB, authdb.NewStore(cfg.Log, cfg.DB))
	authapi.Routes(app, authapi.Config{
		Log:     cfg.Log,
		AuthBus: authbus,
		// Tracer:    cfg.Tracer,
		AppConfig: cfg.AppConfig,
	})
}
