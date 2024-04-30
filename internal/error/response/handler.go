package response

import (
	"net/http"
)

type Handler interface {
	HandleBusinessError(er error, message string, w http.ResponseWriter, r *http.Request)
	HandleIncorrectRequestBodyError(er error, w http.ResponseWriter, r *http.Request)
	HandleIncorrectRequestParamError(er error, w http.ResponseWriter, r *http.Request)
}
