package httpmws

import (
	"fmt"
	"net/http"
)

var allRoutes = make(map[string]bool)
var currentRouter *http.ServeMux

func newRoute(method string, path string, handler http.HandlerFunc, mws ...string) {
	if currentRouter == nil {
		panic("Router not set. Make sure to call UseRouter first")
	}

	newRouteName := method + " " + path

	for _, mw := range mws {
		if !hasMiddleware(mw) {
			panic(fmt.Sprintf("Middleware %s not found for route %s. Make sure to register it first.", mw, newRouteName))
		}
	}

	if allRoutes[newRouteName] {
		panic(fmt.Sprintf("Route %s already registered", newRouteName))
	}

	currentRouter.HandleFunc(newRouteName, useMiddleware(handler, mws...))
	allRoutes[newRouteName] = true

	fmt.Printf("Route %s registered\n", newRouteName)
}

// UseRouter initializes the router
func UseRouter(router *http.ServeMux) *http.ServeMux {
	currentRouter = router
	return currentRouter
}

// RegisterGetRoute creates a new GET route with the given path, http handler and middlewares
func RegisterGetRoute(path string, handler http.HandlerFunc, mws ...string) {
	newRoute(http.MethodGet, path, handler, mws...)
}

// RegisterPostRoute creates a new POST route with the given path, http handler and middlewares
func RegisterPostRoute(path string, handler http.HandlerFunc, mws ...string) {
	newRoute(http.MethodPost, path, handler, mws...)
}

// RegisterPutRoute creates a new PUT route with the given path, http handler and middlewares
func RegisterPutRoute(path string, handler http.HandlerFunc, mws ...string) {
	newRoute(http.MethodPut, path, handler, mws...)
}

// RegisterDeleteRoute creates a new DELETE route with the given path, http handler and middlewares
func RegisterDeleteRoute(path string, handler http.HandlerFunc, mws ...string) {
	newRoute(http.MethodDelete, path, handler, mws...)
}

// RegisterPatchRoute creates a new PATCH route with the given path, http handler and middlewares
func RegisterPatchRoute(path string, handler http.HandlerFunc, mws ...string) {
	newRoute(http.MethodPatch, path, handler, mws...)
}
