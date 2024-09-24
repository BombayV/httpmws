package notfound

import (
	"github.com/bombayv/learning-go/cmd/api/server/router/handlers"
	"net/http"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	handlers.WriteJsonResponse(w, *handlers.NewRouteError(http.StatusNotFound, "Not Found", "The requested resource was not found"))
}
