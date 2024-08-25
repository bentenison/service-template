package web

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/bentenison/microservice/foundation/logger"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) any
type App struct {
	logger *logger.CustomLogger
	build  string
	mux    *mux.Router // change this field to incorporate your own framework
	mw     []MiddlewareFunc
}

func NewApp(log *logger.CustomLogger, build string) *App {
	return &App{
		logger: log,
		build:  build,
		mux:    mux.NewRouter(),
	}
}

// App satisfies the http.Handler interface
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

func (a *App) Handle(path string, handler http.Handler) *mux.Route {
	return a.mux.Handle(path, handler)
}
func wrapMiddleware(mw []MiddlewareFunc, handler HandlerFunc) HandlerFunc {

	// Loop backwards through the middleware invoking each one. Replace the
	// handler with the new wrapped handler. Looping backwards ensures that the
	// first middleware of the slice is the first to be executed by requests.
	for i := len(mw) - 1; i >= 0; i-- {
		mwFunc := mw[i]
		if mwFunc != nil {
			handler = mwFunc.Wrap(handler)
		}
	}

	return handler
}

// Use adds a global middleware to the mux
func (a *App) Use(midFuncs GlobalMiddlewareFunc) {
	a.mux.Use(mux.MiddlewareFunc(midFuncs))
}

// HandleFunc handles a request and also can be used to inject route specific middlewares
func (a *App) HandleFunc(path string, handlerFunc HandlerFunc, midFuncs ...MiddlewareFunc) *mux.Route {
	// tm := mid.NewTransactionMiddleware()
	// h := tm.Wrap(http.HandlerFunc(handlerFunc))
	// h := wrapMiddleware(a.mw, handlerFunc)

	handlerFunc = wrapMiddleware(midFuncs, handlerFunc)
	h := func(w http.ResponseWriter, r *http.Request) {
		// ctx := setTracer(r.Context(), a.tracer)
		// ctx = setWriter(ctx, w)

		// otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(w.Header()))

		res := handlerFunc(w, r)
		if err := Respond(r.Context(), w, res); err != nil {
			a.logger.Error("error responding client", map[string]interface{}{
				"error": err,
			})
		}
	}
	// return a.mux.Handle(path, h)
	return a.mux.HandleFunc(path, h)
}
