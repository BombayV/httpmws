package main

import (
	"fmt"
	"github.com/bombayv/httpmws"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	httpmws.UseRouter(router)

	httpmws.RegisterMw("anon", func(w http.ResponseWriter, r *http.Request) bool {
		if r.Header.Get("Authorization") == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return false
		}
		return true
	})

	httpmws.RegisterGetRoute("/hello", func(w http.ResponseWriter, r *http.Request) {
		type data struct {
			Name string `json:"name"`
		}

		httpmws.JsonResponse(w, http.StatusOK, "OK", &data{Name: "lol"})
	})

	httpmws.RegisterGetRoute("/hello2", func(w http.ResponseWriter, r *http.Request) {
		httpmws.ErrorResponse(w, http.StatusNotFound, "Not found")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("Error starting server: %s", err))

	}
}
