package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	route *router
}

func New() *Engine {
	return &Engine{route: newRouter()}
}

func (e *Engine) ServeHTTP(r http.ResponseWriter, req *http.Request) {
	c := newContext(req, r)
	e.route.handle(c)
}

func (e *Engine) addRoute(method, path string, handlerFunc HandlerFunc) {
	e.route.addRoute(method, path, handlerFunc)
}

func (e *Engine) Get(path string, handlerFunc HandlerFunc) {
	e.addRoute("GET", path, handlerFunc)
}

func (e *Engine) Post(path string, handlerFunc HandlerFunc) {
	e.addRoute("POST", path, handlerFunc)
}

func (e *Engine) Put(path string, handlerFunc HandlerFunc) {
	e.addRoute("PUT", path, handlerFunc)
}

func (e *Engine) Run(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%v", port), e)
}
