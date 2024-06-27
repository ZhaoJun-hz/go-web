package main

import (
	"fmt"
	"github.com/ZhaoJun-hz/go-web/server"
)

func main() {
	engine := server.New()
	userGroup := engine.Group("user")
	userGroup.Get("/list", func(ctx *server.Context) {
		fmt.Fprintf(ctx.W, "user list")
	})
	userGroup.Post("/info", func(ctx *server.Context) {
		fmt.Fprintf(ctx.W, "user info")
	})
	productGroup := engine.Group("product")
	productGroup.Get("/list", func(ctx *server.Context) {
		fmt.Fprintf(ctx.W, "product list")
	})
	productGroup.Post("/info", func(ctx *server.Context) {
		fmt.Fprintf(ctx.W, "product info")
	})
	engine.Run()
}
