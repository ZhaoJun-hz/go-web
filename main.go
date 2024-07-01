package main

import (
	"fmt"
	"github.com/ZhaoJun-hz/go-web/server"
)

func main() {
	engine := server.New()

	userGroup := engine.Group("/user")
	userGroup.Use(func(next server.HandlerFunc) server.HandlerFunc {
		return func(ctx *server.Context) {
			fmt.Println("Router Group Pre Middleware 1")
			next(ctx)
			fmt.Println("Router Group Post Middleware 1")
		}
	})
	userGroup.Use(func(next server.HandlerFunc) server.HandlerFunc {
		return func(ctx *server.Context) {
			fmt.Println("Router Group Pre Middleware 2")
			next(ctx)
			fmt.Println("Router Group Post Middleware 2")
		}
	})

	userGroup.Get("/list", func(ctx *server.Context) {
		fmt.Println("user list handler")
		fmt.Fprintf(ctx.W, "/list")
	}, func(next server.HandlerFunc) server.HandlerFunc {
		return func(ctx *server.Context) {
			fmt.Println("method Pre Middleware 1")
			next(ctx)
			fmt.Println("method Post Middleware 1")
		}
	}, func(next server.HandlerFunc) server.HandlerFunc {
		return func(ctx *server.Context) {
			fmt.Println("method Pre Middleware 2")
			next(ctx)
			fmt.Println("method Post Middleware 2")
		}
	})
	userGroup.Get("/add", func(ctx *server.Context) {
		fmt.Fprintf(ctx.W, "/add")
	})

	productGroup := engine.Group("/product")
	productGroup.Get("/:id/:name", func(ctx *server.Context) {
		fmt.Fprintf(ctx.W, "/:id/:name")
	})

	deptGroup := engine.Group("/dept")
	deptGroup.Get("/*", func(ctx *server.Context) {
		fmt.Fprintf(ctx.W, "/*")
	})

	menuGroup := engine.Group("/menu")
	menuGroup.Get("/*/info", func(ctx *server.Context) {
		fmt.Fprintf(ctx.W, "/*/info")
	})

	engine.Run()
}
