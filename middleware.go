package guren

import (
	"net/http"
)

// Middleware is guren middleware
type Middleware func(ctx *Context, next func()) error

// Middlewares is slice of middleware
type Middlewares []Middleware

func (ms *Middlewares) add(mw Middleware) {
	*ms = append(*ms, mw)
}

func (ms Middlewares) compose() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := &Context{w, r}
		if err := ms[0](ctx, ms.dispatch(ctx, 1)); err != nil {
			panic(err)
		}
	}
}

func (ms Middlewares) dispatch(ctx *Context, i int) func() {
	if i >= len(ms) {
		return func() {}
	}
	return func() {
		if err := ms[i](ctx, ms.dispatch(ctx, i+1)); err != nil {
			panic(err)
		}
	}
}
