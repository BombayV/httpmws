package httpmws

import (
	"fmt"
	"net/http"
)

var AllRoutes = make(map[string]bool)

func newRoute(method string, path string, handler http.HandlerFunc, mws ...string) {
	newRouteName := method + " " + path

	for _, mw := range mws {
		if !hasMiddleware(mw) {
			panic(fmt.Sprintf("Middleware %s not found for route %s. Make sure to register it first.", mw, newRouteName))
		}
	}

	http.HandleFunc(newRouteName, useMiddleware(handler, mws...))

	AllRoutes[newRouteName] = true
}

// NewGetRoute creates a new GET route with the given path, http handler and middlewares
func NewGetRoute(path string, handler http.HandlerFunc, mws ...string) {
	newRoute(http.MethodGet, path, handler, mws...)
}

// NewPostRoute creates a new POST route with the given path, http handler and middlewares
func NewPostRoute(path string, handler http.HandlerFunc, mws ...string) {
	newRoute(http.MethodPost, path, handler, mws...)
}

// NewPutRoute creates a new PUT route with the given path, http handler and middlewares
func NewPutRoute(path string, handler http.HandlerFunc, mws ...string) {
	newRoute(http.MethodPut, path, handler, mws...)
}

// NewDeleteRoute creates a new DELETE route with the given path, http handler and middlewares
func NewDeleteRoute(path string, handler http.HandlerFunc, mws ...string) {
	newRoute(http.MethodDelete, path, handler, mws...)
}

// NewPatchRoute creates a new PATCH route with the given path, http handler and middlewares
func NewPatchRoute(path string, handler http.HandlerFunc, mws ...string) {
	newRoute(http.MethodPatch, path, handler, mws...)
}
