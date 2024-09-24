package logger

import (
	"fmt"
	"net/http"
)

var colors = map[string]string{
	"GET":      "\033[1;32m",
	"POST":     "\033[1;34m",
	"PUT":      "\033[1;33m",
	"DELETE":   "\033[1;31m",
	"NOTFOUND": "\033[1;93m",
}

var reset = "\033[0m"

func LogRequest(r *http.Request) {
	method := r.Method
	ip := r.RemoteAddr
	path := r.URL.Path
	color := colors[method]

	log := color + method + reset + " " + path + " " + ip
	fmt.Println(log)
}

func LogRequestError(error string, r *http.Request) {
	color := colors["DELETE"]
	ip := r.RemoteAddr
	method := r.Method
	path := r.URL.Path
	log := color + method + reset + " " + path + " " + ip + " " + error
	fmt.Println(log)
}

func LogNotFound(r *http.Request) {
	ip := r.RemoteAddr
	path := r.URL.Path
	color := colors["NOTFOUND"]
	log := color + "NOT FOUND" + reset + " " + path + " " + ip
	fmt.Println(log)
}
