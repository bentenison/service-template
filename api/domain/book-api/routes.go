package bookapi

import (
	"github.com/bentenison/microservice/app/domain/bookapp"
	"github.com/bentenison/microservice/app/sdk/mid"
	"github.com/bentenison/microservice/business/domain/bookbus"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/web"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Config struct {
	Log     *logger.CustomLogger
	BookBus *bookbus.Business
	Tracer  *trace.TracerProvider
}

func Routes(app *web.App, cfg Config) {
	api := newAPI(bookapp.NewApp(cfg.BookBus, cfg.Tracer))
	app.Use(mid.TraceIdMiddleware())
	app.Handle("GET", "/test", api.query)
}
