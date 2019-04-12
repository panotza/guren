package guren

import (
	"net/http"
)

type (
	// Guren is the top-level framework instance.
	Guren struct {
		server *http.Server
		ms     Middleware
	}
)

// New creates new Guren
func New() *Guren {
	return &Guren{}
}

// Use uses the given middleware
func (g *Guren) Use(m Middleware) *Guren {
	if g.ms == nil {
		g.ms = m
	} else {
		g.ms = g.ms.add(m)
	}
	return g
}

// Listen start server at given port
func (g *Guren) Listen(addr string) error {
	g.server.Addr = addr
	g.server.Handler = g.ms
	return g.server.ListenAndServe()
}
