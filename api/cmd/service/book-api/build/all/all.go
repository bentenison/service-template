package all

import (
	bookapi "github.com/bentenison/microservice/api/domain/book-api"
	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/domain/bookbus"
	"github.com/bentenison/microservice/business/domain/bookbus/stores/bookdb"
	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/foundation/web"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg mux.Config) {
	delegate := delegate.New(cfg.Log)
	bookbus := bookbus.NewBusiness(cfg.Log, cfg.DB.SQL, delegate, bookdb.NewStore(cfg.Log, cfg.DB.SQL))
	// Construct the business domain packages we need here so we are using the
	// sames instances for the different set of domain apis.
	bookapi.Routes(app, bookapi.Config{
		Log:     cfg.Log,
		BookBus: bookbus,
		// Tracer:  cfg.Tracer,
	})

}
