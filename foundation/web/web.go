package web

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/bentenison/microservice/foundation/logger"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)
type App struct {
	logger *logger.CustomLogger
	build  string
	mux    *mux.Router // change this field to incorporate your own framework
}

func NewApp(log *logger.CustomLogger, build string) *App {
	return &App{
		logger: log,
		build:  build,
		mux:    mux.NewRouter(),
	}
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

func (a *App) Handle(path string, handler http.Handler) *mux.Route {
	return a.mux.Handle(path, handler)
}
func (a *App) HandleFunc(path string, handleFunc HandleFunc) *mux.Route {
	return a.mux.HandleFunc(path, handleFunc)
}
