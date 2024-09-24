package router

import (
	"github.com/bombayv/learning-go/cmd/api/server/router/middleware"
	"net/http"
)

func NewRoute(method string, path string, handler http.HandlerFunc, mws ...string) {
	var mw []string
	hasAnon := false
	for _, m := range mws {
		if m == "anonymous" {
			hasAnon = true
			break
		}

		mw = append(mw, m)
	}

	if !hasAnon {
		mw = append(mw, "anonymous")
	}

	http.HandleFunc(method+" "+path, middleware.UseMiddleware(handler, mw...))
	middleware.RegisterRoute(method, path)
}

func NewGetRoute(path string, handler http.HandlerFunc, mws ...string) {
	NewRoute(http.MethodGet, path, handler, mws...)
}

func NewPostRoute(path string, handler http.HandlerFunc, mws ...string) {
	NewRoute(http.MethodPost, path, handler, mws...)
}

func NewPutRoute(path string, handler http.HandlerFunc, mws ...string) {
	NewRoute(http.MethodPut, path, handler, mws...)
}

func NewDeleteRoute(path string, handler http.HandlerFunc, mws ...string) {
	NewRoute(http.MethodDelete, path, handler, mws...)
}
