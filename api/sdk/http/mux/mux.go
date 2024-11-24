package mux

import (
	"github.com/bentenison/microservice/foundation/conf"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/web"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel/sdk/trace"
)

type RouteAdder interface {
	Add(app *web.App, cfg Config)
}
type DataSource struct {
	SQL *sqlx.DB
	MGO *mongo.Database
	RDB *redis.Client
}
type Config struct {
	Build string
	Log   *logger.CustomLogger
	// Auth       *auth.Auth
	// AuthClient *authclient.Client
	DB        DataSource
	Tracer    *trace.TracerProvider
	AppConfig *conf.Config
}

func WebAPI(cfg Config, routeAdder RouteAdder) *web.App {
	app := web.NewApp(cfg.Log, cfg.Build, cfg.AppConfig)
	routeAdder.Add(app, cfg)
	return app
}
