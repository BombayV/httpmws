package httpmws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var colors = map[string]string{
	"red":      "\033[31m",
	"green":    "\033[32m",
	"yellow":   "\033[33m",
	"blue":     "\033[34m",
	"purple":   "\033[35m",
	"cyan":     "\033[36m",
	"white":    "\033[37m",
	"reset":    "\033[0m",
	"bgred":    "\033[41m",
	"bggreen":  "\033[42m",
	"bgyellow": "\033[43m",
	"bgblue":   "\033[44m",
	"bgpurple": "\033[45m",
	"bgcyan":   "\033[46m",
	"bgwhite":  "\033[47m",
}

var methodToColor = map[string]string{
	http.MethodGet:    colors["green"],
	http.MethodPost:   colors["blue"],
	http.MethodPut:    colors["purple"],
	http.MethodDelete: colors["red"],
	http.MethodPatch:  colors["cyan"],
}

type errorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type jsonReponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func statusCodeToColor(statusCode int) string {
	switch {
	case statusCode < 200:
		return colors["reset"]
	case statusCode < 300:
		return colors["green"]
	case statusCode < 400:
		return colors["cyan"]
	case statusCode < 500:
		return colors["yellow"]
	default:
		return colors["red"]
	}
}

func generateLogMessage(r *http.Request, statusCode int, duration time.Duration) string {
	title := fmt.Sprintf("%s%s%s", colors["blue"], "[httpmws]", colors["reset"])
	code := fmt.Sprintf("%s%d%s", statusCodeToColor(statusCode), statusCode, colors["reset"])
	method := fmt.Sprintf("%s%s%s", methodToColor[r.Method], r.Method, colors["reset"])

	return fmt.Sprintf("%s | %s | %s | %s | %s", title, code, duration, method, r.URL.Path)
}

func ErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	errResponse := &errorResponse{}
	errResponse.Error.Code = code
	errResponse.Error.Message = message

	jData, err := json.Marshal(errResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	written, err := w.Write(jData)
	if err != nil {
		panic(fmt.Sprintf("Error writing response: %s", err))
	}

	fmt.Printf("Response written: %d bytes\n", written)
}

func JsonResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonResponse := &jsonReponse{Code: code, Message: message, Data: data}

	jData, err := json.Marshal(jsonResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	written, err := w.Write(jData)
	if err != nil {
		panic(fmt.Sprintf("Error writing response: %s", err))
	}

	fmt.Printf("Response written: %d bytes\n", written)
}
