package server

import (
	"log"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type routerGroup struct {
	name           string
	handlerFuncMap map[string]HandlerFunc
}

func (r *routerGroup) Add(name string, handlerFunc HandlerFunc) {
	r.handlerFuncMap[name] = handlerFunc
}

func (r *router) Group(name string) *routerGroup {
	rg := &routerGroup{
		name:           name,
		handlerFuncMap: make(map[string]HandlerFunc),
	}
	r.routerGroups = append(r.routerGroups, rg)
	return rg
}

type router struct {
	routerGroups []*routerGroup
}

type Engine struct {
	router
}

func New() *Engine {
	return &Engine{
		router: router{},
	}
}

func (e *Engine) Run() {
	for _, rg := range e.routerGroups {
		for name, handlerFunc := range rg.handlerFuncMap {
			http.HandleFunc("/"+rg.name+name, handlerFunc)
		}
	}

	err := http.ListenAndServe(":8111", nil)
	if err != nil {
		log.Fatal(err)
	}
}
