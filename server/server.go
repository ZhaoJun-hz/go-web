package server

import (
	"log"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type routerGroup struct {
	name             string
	handlerFuncMap   map[string]HandlerFunc
	handlerMethodMap map[string][]string
}

func (r *routerGroup) Get(name string, handlerFunc HandlerFunc) {
	r.handlerFuncMap[name] = handlerFunc
	r.handlerMethodMap[http.MethodGet] = append(r.handlerMethodMap[http.MethodGet], name)
}

func (r *routerGroup) POST(name string, handlerFunc HandlerFunc) {
	r.handlerFuncMap[name] = handlerFunc
	r.handlerMethodMap[http.MethodPost] = append(r.handlerMethodMap[http.MethodPost], name)
}

func (r *routerGroup) PUT(name string, handlerFunc HandlerFunc) {
	r.handlerFuncMap[name] = handlerFunc
	r.handlerMethodMap[http.MethodPut] = append(r.handlerMethodMap[http.MethodPut], name)
}

func (r *routerGroup) DELETE(name string, handlerFunc HandlerFunc) {
	r.handlerFuncMap[name] = handlerFunc
	r.handlerMethodMap[http.MethodDelete] = append(r.handlerMethodMap[http.MethodDelete], name)
}

func (r *router) Group(name string) *routerGroup {
	rg := &routerGroup{
		name:             name,
		handlerFuncMap:   make(map[string]HandlerFunc),
		handlerMethodMap: make(map[string][]string),
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

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	for _, rg := range e.routerGroups {
		for name, handlerFunc := range rg.handlerFuncMap {
			url := "/" + rg.name + name
			// 路由是匹配的
			if r.RequestURI == url {
				// 匹配Method
				routers, ok := rg.handlerMethodMap[method]
				if ok {
					// method 里面有路由
					for _, routerName := range routers {
						tempUrl := "/" + rg.name + routerName
						if url == tempUrl {
							handlerFunc(w, r)
							return
						}
					}
				}
				// 方法不匹配
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func (e *Engine) Run() {
	//for _, rg := range e.routerGroups {
	//	for name, handlerFunc := range rg.handlerFuncMap {
	//		http.HandleFunc("/"+rg.name+name, handlerFunc)
	//	}
	//}
	http.Handle("/", e)
	err := http.ListenAndServe(":8111", nil)
	if err != nil {
		log.Fatal(err)
	}
}
