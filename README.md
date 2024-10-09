# httpmws
A simple HTTP server using golang and net/http package, with middleware support. Mainly used for private projects.

## Description
This package is just a simple wrapper around the net/http package, with middleware support. It allows you to easily create a server with middleware support, and register routes with middleware.

## Installation
```bash
go get github.com/bombayv/httpmws
```

## Usage
```golang
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

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("Error starting server: %s", err))

	}
}
```

## License
MIT
```

