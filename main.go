package main

import (
	"fmt"
	"github.com/ZhaoJun-hz/go-web/server"
	"net/http"
)

func main() {
	engine := server.New()
	userGroup := engine.Group("user")
	userGroup.Add("/list", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "user list")
	})
	userGroup.Add("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "user info")
	})
	productGroup := engine.Group("product")
	productGroup.Add("/list", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "product list")
	})
	productGroup.Add("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "product info")
	})
	engine.Run()
}
