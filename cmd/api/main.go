package main

import (
	"fmt"
	"github.com/bombayv/learning-go/cmd/api/server/router/hello"
	"github.com/bombayv/learning-go/cmd/api/server/router/middleware"
	"github.com/bombayv/learning-go/cmd/api/server/router/notfound"
	"net/http"
)

func main() {
	middleware.InitMiddlewares()
	notfound.InitRoutes()
	hello.InitRoutes()

	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}
}
