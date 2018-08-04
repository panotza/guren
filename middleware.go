package guren

import (
	"net/http"
)

// Middleware is guren middleware
type Middleware func(ctx *Context, next func()) error

func (m Middleware) add(mw Middleware) Middleware {
	return func(ctx *Context, next func()) error {
		return m(ctx, func() {
			mw(ctx, next)
		})
	}
}

func noop() {}

func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m(&Context{w, r}, noop)
}
