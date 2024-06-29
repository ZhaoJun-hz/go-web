package main

import (
	"fmt"
	"github.com/ZhaoJun-hz/go-web/server"
)

func main() {
	engine := server.New()
	userGroup := engine.Group("user")
	userGroup.Get("/list", func(ctx *server.Context) {
		fmt.Fprintf(ctx.W, "/list")
	})
	userGroup.Get("/list/hello", func(ctx *server.Context) {
		fmt.Fprintf(ctx.W, "/list/hello")
	})
	userGroup.Get("/list/*/hello", func(ctx *server.Context) {
		fmt.Fprintf(ctx.W, "/list/*/hello")
	})
	userGroup.Post("/info", func(ctx *server.Context) {
		fmt.Fprintf(ctx.W, "/info")
	})
	productGroup := engine.Group("product")
	productGroup.Get("/list", func(ctx *server.Context) {
		fmt.Fprintf(ctx.W, "/list")
	})
	productGroup.Get("/hello/get", func(ctx *server.Context) {
		fmt.Fprintf(ctx.W, "/hello/get")
	})
	productGroup.Post("/info", func(ctx *server.Context) {
		fmt.Fprintf(ctx.W, "/info")
	})
	productGroup.Get("/get/:id", func(ctx *server.Context) {
		fmt.Fprintf(ctx.W, "/get/:id")
	})
	engine.Run()
}
