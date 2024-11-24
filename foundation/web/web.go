package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bentenison/microservice/foundation/conf"
	"github.com/bentenison/microservice/foundation/logger"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) any
type App struct {
	logger    *logger.CustomLogger
	build     string
	mux       *gin.Engine
	appConfig *conf.Config
}

func NewApp(log *logger.CustomLogger, build string, cfg *conf.Config) *App {
	return &App{
		logger:    log,
		build:     build,
		mux:       gin.Default(),
		appConfig: cfg,
	}
}
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

func (a *App) Handle(method, route string, handler gin.HandlerFunc) {
	switch method {
	case "GET":
		a.mux.GET(route, handler)
	case "POST":
		a.mux.POST(route, handler)
	case "PUT":
		a.mux.PUT(route, handler)
	case "DELETE":
		a.mux.DELETE(route, handler)
	// Add more methods as needed
	default:
		panic("unsupported method")
	}
}

func (a *App) Run(addr string) error {
	return a.mux.Run(addr)
}
func (a *App) Use(mids ...gin.HandlerFunc) {
	for _, midFunc := range mids {
		a.mux.Use(midFunc)
	}
}

// func (a *App) Run(addr string) error {
// 	return a.mux.
// }
