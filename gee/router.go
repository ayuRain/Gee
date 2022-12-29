package gee

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: map[string]HandlerFunc{}}
}

func (r *router) addRoute(method, path string, handlerFunc HandlerFunc) {
	key := method + "-" + path
	if _, ok := r.handlers[key]; ok {
		panic("repeat router!")
	}
	r.handlers[key] = handlerFunc
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handle, ok := r.handlers[key]; ok {
		handle(c)
	} else {
		c.String(404, "error")
	}
}
