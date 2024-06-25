package server

import (
	"log"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type router struct {
	handlerFuncMap map[string]HandlerFunc
}

func (r *router) Add(name string, handlerFunc HandlerFunc) {
	r.handlerFuncMap[name] = handlerFunc
}

type Engine struct {
	router
}

func New() *Engine {
	return &Engine{
		router: router{
			handlerFuncMap: make(map[string]HandlerFunc),
		},
	}
}

func (e *Engine) Run() {
	for name, handlerFunc := range e.handlerFuncMap {
		http.HandleFunc(name, handlerFunc)
	}
	err := http.ListenAndServe(":8111", nil)
	if err != nil {
		log.Fatal(err)
	}
}
