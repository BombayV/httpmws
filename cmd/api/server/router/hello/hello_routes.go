package hello

import "net/http"

func getHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
