package httpmws

//
//var colors = map[string]string{
//	"GET":    "\033[1;32m",
//	"POST":   "\033[1;34m",
//	"PUT":    "\033[1;33m",
//	"DELETE": "\033[1;31m",
//}
//
//var reset = "\033[0m"
//
//func logRequest(r *http.Request) {
//	method := r.Method
//	ip := r.RemoteAddr
//	path := r.URL.Path
//	color := colors[method]
//
//	log := color + method + reset + " " + path + " " + ip
//	fmt.Println(log)
//}
//
//func logRequestError(error string, r *http.Request) {
//	color := colors["DELETE"]
//	ip := r.RemoteAddr
//	method := r.Method
//	path := r.URL.Path
//	log := color + method + reset + " " + path + " " + ip + " " + error
//	fmt.Println(log)
//}

type Logger interface {
}
