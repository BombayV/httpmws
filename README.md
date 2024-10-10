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
```
MIT License

Copyright (c) 2024 Mauricio Rivera

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

