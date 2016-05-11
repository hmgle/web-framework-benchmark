package main

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/vishr/web-framework-benchmark/common"
)

// Test 1: Dynamic Route
func dynamicRoute(c *iris.Context) {
	// c.Write([]byte(fmt.Sprintf("team: %s, user: %s", chi.URLParam(ctx, "id"), chi.URLParam(ctx, "user"))))
	c.Write(fmt.Sprintf("team: %s, user: %s", c.Param("id"), c.Param("user")))
}

func main() {
	for _, r := range common.DynamicRoutes {
		switch r.Method {
		case "GET":
			iris.Get(r.Path, dynamicRoute)
		case "POST":
			iris.Post(r.Path, dynamicRoute)
		case "PUT":
			iris.Put(r.Path, dynamicRoute)
		case "PATCH":
			iris.Patch(r.Path, dynamicRoute)
		case "DELETE":
			iris.Delete(r.Path, dynamicRoute)
		default:
			panic("method not supported")
		}
	}
	iris.Listen(":8080")
}
