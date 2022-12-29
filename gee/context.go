package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MapInterface map[string]interface{}

type Context struct {
	// Http request and response obj
	Req   *http.Request
	Write http.ResponseWriter
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
}

func newContext(req *http.Request, r http.ResponseWriter) *Context {
	return &Context{
		Req:    req,
		Write:  r,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) SetHeader(key string, value string) {
	c.Write.Header().Set(key, value)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Write.WriteHeader(code)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/plain")
	c.Write.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/plain")
	c.Status(code)
	encoder := json.NewEncoder(c.Write)
	err := encoder.Encode(obj)
	if err != nil {
		http.Error(c.Write, err.Error(), 500)
	}
}

func (c *Context) HTML(code int, html string) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/html")
	c.Write.Write([]byte(html))
}

func (c *Context) DATA(code int, data []byte) {
	c.Status(code)
	c.Write.Write(data)
}
