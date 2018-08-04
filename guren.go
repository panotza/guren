package guren

import (
	"net/http"
)

type (
	// Guren is the top-level framework instance.
	Guren struct {
		server      *http.Server
		middlewares Middlewares
	}
)

// New creates new Guren
func New() *Guren {
	return &Guren{}
}

// Use uses the given middleware
func (g *Guren) Use(m Middleware) *Guren {
	g.middlewares.add(m)
	return g
}

// Listen start server at given port
func (g *Guren) Listen(addr string) {
	err := http.ListenAndServe(addr, g.middlewares)
	if err != nil {
		panic(err)
	}
}
