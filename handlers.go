package httpmws

type jsonReponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func HttpMwsError(code int, message string) *jsonReponse {
	return &jsonReponse{Code: code, Message: message}
}

func HttpMwsJsonResponse() {

}
