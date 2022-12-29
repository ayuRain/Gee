package main

import (
	"Gee/gee"
)

func main() {
	r := gee.New()
	r.Get("/", func(c *gee.Context) {
		c.String(200, "hello")
	})
	r.Run(9999)
}
