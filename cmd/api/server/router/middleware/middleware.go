package middleware

import (
	"net/http"
)

var allMiddleware = make(map[string]func(w http.ResponseWriter, r *http.Request) bool)
var allRoutes = make(map[string]bool)

// getMiddleware returns the middleware function
// with the given name
func getMiddleware(name string) func(w http.ResponseWriter, r *http.Request) bool {
	return allMiddleware[name]
}

// RegisterMiddleware registers a new middleware
// with the given name and function
func registerHandler(name string, fn func(w http.ResponseWriter, r *http.Request) bool) {
	allMiddleware[name] = fn
}

// RegisterRoute registers a new route
// with the given method and path.
// Mainly used on router.go
func RegisterRoute(method string, path string) {
	allRoutes[method+"_"+path] = true
}

// InitMiddlewares initializes all middlewares
// that are available in the application
func InitMiddlewares() {
	registerHandler("auth", auth)
	registerHandler("anonymous", anonymous)
}

// UseMiddleware is a helper function that takes a handler
// and a list of middlewares to be executed before the handler
func UseMiddleware(next http.HandlerFunc, mws ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, m := range mws {
			if !getMiddleware(m)(w, r) {
				return
			}
		}

		next(w, r)
	}
}
