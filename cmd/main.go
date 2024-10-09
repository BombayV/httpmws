package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	router.HandleFunc("test/{id}", func(w http.ResponseWriter, r *http.Request) {
		a := r.PathValue("id")
		w.Write([]byte("Hello World " + a))
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
