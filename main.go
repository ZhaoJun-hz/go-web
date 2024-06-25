package main

import (
	"fmt"
	"github.com/ZhaoJun-hz/go-web/server"
	"net/http"
)

func main() {
	engine := server.New()
	engine.Add("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})
	engine.Run()
}
