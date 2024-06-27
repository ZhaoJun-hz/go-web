package main

import (
	"fmt"
	"github.com/ZhaoJun-hz/go-web/server"
	"net/http"
)

func main() {
	engine := server.New()
	userGroup := engine.Group("user")
	userGroup.Get("/list", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "user list")
	})
	userGroup.POST("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "user info")
	})
	productGroup := engine.Group("product")
	productGroup.Get("/list", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "product list")
	})
	productGroup.POST("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "product info")
	})
	engine.Run()
}
