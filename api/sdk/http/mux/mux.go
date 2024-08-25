package mux

import (
	"net/http"

	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/web"
	"github.com/jmoiron/sqlx"
)

type RouteAdder interface {
	Add(app *web.App, cfg Config)
}
type Config struct {
	Build string
	Log   *logger.CustomLogger
	// Auth       *auth.Auth
	// AuthClient *authclient.Client
	DB *sqlx.DB
	// Tracer     trace.Tracer
}

func WebAPI(cfg Config, routeAdder RouteAdder) http.Handler {
	app := web.NewApp(cfg.Log, cfg.Build)
	routeAdder.Add(app, cfg)
	return app
}
