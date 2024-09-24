package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RouteError struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func NewRouteError(code int, error string, message string) *RouteError {
	return &RouteError{
		Code:    code,
		Error:   error,
		Message: message,
	}
}

// WriteJsonResponse writes the response to the client
// and sets the appropriate headers
func WriteJsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if data.(RouteError).Code != 0 {
		w.WriteHeader(data.(RouteError).Code)
	}

	jData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling response: ", err)
		return
	}

	write, err := w.Write(jData)
	if err != nil {
		fmt.Println("Error writing response: ", err)
		return
	}

	if write == 0 {
		fmt.Println("No bytes written")
	}
}
