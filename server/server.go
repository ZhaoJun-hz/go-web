package server

import (
	"log"
	"net/http"
)

type HandlerFunc func(ctx *Context)

type routerGroup struct {
	name string
	// key 路由path map[string]HandlerFunc，key，请求类型，GET、POST等
	handlerFuncMap map[string]map[string]HandlerFunc
	tree           *treeNode
}

func (r routerGroup) handle(name string, method string, handlerFunc HandlerFunc) {
	_, ok := r.handlerFuncMap[name]
	if !ok {
		r.handlerFuncMap[name] = make(map[string]HandlerFunc)
	}
	_, ok = r.handlerFuncMap[name][method]
	if ok {
		panic("有重复的路由")
	}
	r.handlerFuncMap[name][http.MethodGet] = handlerFunc

	r.tree.Put(name)
}

func (r *routerGroup) Get(name string, handlerFunc HandlerFunc) {
	r.handle(name, http.MethodGet, handlerFunc)
}

func (r *routerGroup) Post(name string, handlerFunc HandlerFunc) {
	r.handle(name, http.MethodPost, handlerFunc)
}

func (r *routerGroup) Put(name string, handlerFunc HandlerFunc) {
	r.handle(name, http.MethodPut, handlerFunc)
}

func (r *routerGroup) Delete(name string, handlerFunc HandlerFunc) {
	r.handle(name, http.MethodDelete, handlerFunc)
}

func (r *router) Group(name string) *routerGroup {
	rg := &routerGroup{
		name:           name,
		handlerFuncMap: make(map[string]map[string]HandlerFunc),
		tree:           &treeNode{name: "/", isEnd: true, children: make([]*treeNode, 0)},
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
		routerPath := SubStringLast(r.RequestURI, "/"+rg.name)
		node := rg.tree.Get(routerPath)
		if node != nil {
			// 路由匹配上了
			ctx := &Context{
				W: w,
				R: r,
			}
			handler, ok := rg.handlerFuncMap[node.routerName][method]
			if ok {
				handler(ctx)
				return
			}
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func (e *Engine) Run() {
	http.Handle("/", e)
	err := http.ListenAndServe(":8111", nil)
	if err != nil {
		log.Fatal(err)
	}
}
