package main

import (
	"log"

	"github.com/panotza/guren"
)

func main() {
	g := guren.New()
	g.Use(func(ctx *guren.Context, next func()) error {
		log.Println("down stream 1")
		next()
		log.Println("up stream 1")
		ctx.W.Write([]byte("hello world"))
		return nil
	})
	g.Use(func(ctx *guren.Context, next func()) error {
		log.Println("down stream 2")
		next()
		log.Println("up stream 2")
		return nil
	})
	g.Listen(":8080")
}
