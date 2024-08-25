package bookapi

import (
	"github.com/bentenison/microservice/app/domain/bookapp"
	"github.com/bentenison/microservice/app/sdk/mid"
	"github.com/bentenison/microservice/business/domain/bookbus"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/web"
)

type Config struct {
	Log     *logger.CustomLogger
	BookBus *bookbus.Business
}

func Routes(app *web.App, cfg Config) {
	api := newAPI(bookapp.NewApp(cfg.BookBus))
	app.HandleFunc("/test", api.query, mid.TraceTdMiddleware).Methods("GET")
}
