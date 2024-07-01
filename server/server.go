package server

import (
	"log"
	"net/http"
)

type HandlerFunc func(ctx *Context)

// 中间件逻辑
type MiddlewareFunc func(HandlerFunc) HandlerFunc

type routerGroup struct {
	name string
	// key 路由path map[string]HandlerFunc，key，请求类型，GET、POST等
	handlerFuncMap           map[string]map[string]HandlerFunc
	middlewareHandlerFuncMap map[string]map[string][]MiddlewareFunc
	tree                     *treeNode
	// 中间件
	middlewares []MiddlewareFunc
}

func (r *routerGroup) Use(middlewares ...MiddlewareFunc) {
	r.middlewares = append(r.middlewares, middlewares...)
}

func (r routerGroup) handle(name string, method string, handlerFunc HandlerFunc, middlewares ...MiddlewareFunc) {
	_, ok := r.handlerFuncMap[name]
	if !ok {
		r.handlerFuncMap[name] = make(map[string]HandlerFunc)
		r.middlewareHandlerFuncMap[name] = make(map[string][]MiddlewareFunc)
	}
	_, ok = r.handlerFuncMap[name][method]
	if ok {
		panic("有重复的路由")
	}
	r.handlerFuncMap[name][method] = handlerFunc
	r.middlewareHandlerFuncMap[name][method] = append(r.middlewareHandlerFuncMap[name][method], middlewares...)
	r.tree.Put(name)
}

func (r *routerGroup) Get(name string, handlerFunc HandlerFunc, middlewares ...MiddlewareFunc) {
	r.handle(name, http.MethodGet, handlerFunc, middlewares...)
}

func (r *routerGroup) Post(name string, handlerFunc HandlerFunc, middlewares ...MiddlewareFunc) {
	r.handle(name, http.MethodPost, handlerFunc, middlewares...)
}

func (r *routerGroup) Put(name string, handlerFunc HandlerFunc, middlewares ...MiddlewareFunc) {
	r.handle(name, http.MethodPut, handlerFunc, middlewares...)
}

func (r *routerGroup) Delete(name string, handlerFunc HandlerFunc, middlewares ...MiddlewareFunc) {
	r.handle(name, http.MethodDelete, handlerFunc, middlewares...)
}

// 执行handler 方法
func (r *routerGroup) methodHandler(routerPath string, method string, handler HandlerFunc, ctx *Context) {
	// 路由组级别中间件
	for _, middleware := range r.middlewares {
		handler = middleware(handler)
	}
	// 单个路由中间件
	middlewareFuncs := r.middlewareHandlerFuncMap[routerPath][method]
	for _, middlewareFunc := range middlewareFuncs {
		handler = middlewareFunc(handler)
	}
	handler(ctx)
}

func (r *router) Group(name string) *routerGroup {
	rg := &routerGroup{
		name:                     name,
		handlerFuncMap:           make(map[string]map[string]HandlerFunc),
		middlewareHandlerFuncMap: make(map[string]map[string][]MiddlewareFunc),
		tree:                     &treeNode{name: "/", isEnd: true, children: make([]*treeNode, 0)},
		middlewares:              make([]MiddlewareFunc, 0),
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
	e.httpRequestHandler(w, r)
}

func (e *Engine) httpRequestHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	for _, rg := range e.routerGroups {
		routerPath := SubStringLast(r.RequestURI, rg.name)
		node := rg.tree.Get(routerPath)
		if node != nil {
			// 路由匹配上了
			ctx := &Context{
				W: w,
				R: r,
			}
			handler, ok := rg.handlerFuncMap[node.routerName][method]
			if ok {
				rg.methodHandler(routerPath, method, handler, ctx)
				//handler(ctx)
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
