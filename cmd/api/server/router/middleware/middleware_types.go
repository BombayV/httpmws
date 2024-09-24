package middleware

import (
	"github.com/bombayv/learning-go/cmd/api/server/router/handlers"
	"github.com/bombayv/learning-go/cmd/api/server/router/logger"
	"net/http"
)

var authUser = "admin"
var authPass = "admin"
var authJWT = "secret"

func anonymous(w http.ResponseWriter, r *http.Request) bool {
	if _, ok := allRoutes[r.Method+"_"+r.URL.Path]; !ok {
		logger.LogNotFound(r)
	} else {
		logger.LogRequest(r)
	}

	return true
}

func auth(w http.ResponseWriter, r *http.Request) bool {
	logger.LogRequestError("Not authorized", r)

	handlers.WriteJsonResponse(w, *handlers.NewRouteError(http.StatusUnauthorized, "Not authorized", "You are not authorized to access this resource"))
	return false
}
