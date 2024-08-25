package web

import "net/http"

type Middleware interface {
	Wrap(http.Handler) http.Handler
}

//GlobalMiddlewareFunc is a global middleware func which is added to every route
type GlobalMiddlewareFunc func(http.Handler) http.Handler

// MiddlewareFunc is an adapter to allow the use of ordinary functions as Middleware.
type MiddlewareFunc func(next http.Handler) HandlerFunc

// Wrap calls the function itself to fulfill the Middleware interface.
func (m MiddlewareFunc) Wrap(next http.Handler) HandlerFunc {
	return m(next)
}
