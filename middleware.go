package httpmws

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

var allMiddleware = make(map[string]func(w http.ResponseWriter, r *http.Request) bool)

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func getMiddleware(name string) func(w http.ResponseWriter, r *http.Request) bool {
	return allMiddleware[name]
}

func hasMiddleware(name string) bool {
	_, ok := allMiddleware[name]
	return ok
}

func useMiddleware(next http.HandlerFunc, mws ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrappedWriter{w, http.StatusOK}
		for _, m := range mws {
			if !getMiddleware(m)(wrapped, r) {
				log.Printf(generateLogMessage(r, wrapped.statusCode, time.Since(start)))
				return
			}
		}

		next.ServeHTTP(wrapped, r)
		log.Printf(generateLogMessage(r, wrapped.statusCode, time.Since(start)))
	}
}

// RegisterMw registers a new middleware
// with the given name and function
func RegisterMw(name string, fn func(w http.ResponseWriter, r *http.Request) bool) {
	if hasMiddleware(name) {
		panic(fmt.Sprintf("Middleware %s already registered", name))
	}

	allMiddleware[name] = fn
}
