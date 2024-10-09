package httpmws

import (
	"fmt"
	"net/http"
)

var AllMiddleware = make(map[string]func(w http.ResponseWriter, r *http.Request) bool)

func getMiddleware(name string) func(w http.ResponseWriter, r *http.Request) bool {
	return AllMiddleware[name]
}

func hasMiddleware(name string) bool {
	_, ok := AllMiddleware[name]
	return ok
}

func useMiddleware(next http.HandlerFunc, mws ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, m := range mws {
			if !getMiddleware(m)(w, r) {
				return
			}
		}

		next(w, r)
	}
}

// RegisterMw registers a new middleware
// with the given name and function
func RegisterMw(name string, fn func(w http.ResponseWriter, r *http.Request) bool) {
	if hasMiddleware(name) {
		panic(fmt.Sprintf("Middleware %s already registered", name))
	}

	AllMiddleware[name] = fn
}
