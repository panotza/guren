package guren

import "net/http"

// M converts http Handler to guren middleware
func M(h http.Handler) Middleware {
	return func(ctx *Context, next func()) error {
		h.ServeHTTP(ctx.W, ctx.R)
		next()
		return nil
	}
}
