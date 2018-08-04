package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/panotza/guren"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	g := guren.New()

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/hello/:name", hello)

	g.Use(func(ctx *guren.Context, next func()) error {
		log.Println("down stream")
		next()
		log.Println("up stream")
		return nil
	})
	g.Use(guren.M(router))
	g.Listen(":8080")
}
