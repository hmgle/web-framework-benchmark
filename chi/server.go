package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	"github.com/hmgle/chi"

	"github.com/valyala/fasthttp"
	"github.com/vishr/web-framework-benchmark/common"
)

// Test 1: Dynamic Route
func dynamicRoute(ctx context.Context, fctx *fasthttp.RequestCtx) {
	fctx.Write([]byte(fmt.Sprintf("team: %s, user: %s", chi.URLParam(ctx, "id"), chi.URLParam(ctx, "user"))))
}

func main() {
	router := chi.NewRouter()

	for _, r := range common.DynamicRoutes {
		switch r.Method {
		case "GET":
			router.Get(r.Path, dynamicRoute)
		case "POST":
			router.Post(r.Path, dynamicRoute)
		case "PUT":
			router.Put(r.Path, dynamicRoute)
		case "PATCH":
			router.Patch(r.Path, dynamicRoute)
		case "DELETE":
			router.Delete(r.Path, dynamicRoute)
		default:
			panic("method not supported")
		}
	}

	log.Fatal(fasthttp.ListenAndServe(":8080", router.ServeHTTP))
}
