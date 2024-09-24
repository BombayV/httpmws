package notfound

import (
	"github.com/bombayv/learning-go/cmd/api/server/router"
)

func InitRoutes() {
	router.NewGetRoute("/*", notFound)
}
