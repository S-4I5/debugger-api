package response

import (
	error2 "debugger-api/internal/error"
	"github.com/go-chi/render"
	"net/http"
)

type Handler struct {
	error2.Source
}

func NewErrorResponseHandler(source error2.Source) *Handler {
	return &Handler{source}
}

func (h *Handler) HandleBusinessError(er error, errorMessage string, w http.ResponseWriter, r *http.Request) {

	render.Status(r, 400)
	render.JSON(w, r,
		error2.ResponseDto{
			Status:     "400",
			Error:      errorMessage,
			Message:    h.Source.GetMessage(errorMessage),
			StackTrace: er.Error(),
		},
	)

	return
}

func (h *Handler) HandleIncorrectRequestBodyError(er error, w http.ResponseWriter, r *http.Request) {

	errorMessage := "api.request.body.incorrect"

	render.Status(r, 400)
	render.JSON(w, r,
		error2.ResponseDto{
			Status:     "420",
			Error:      errorMessage,
			Message:    h.Source.GetMessage(errorMessage),
			StackTrace: er.Error(),
		},
	)

	return
}

func (h *Handler) HandleIncorrectRequestParamError(er error, w http.ResponseWriter, r *http.Request) {

	errorMessage := "api.request.param.incorrect"

	render.Status(r, 400)
	render.JSON(w, r,
		error2.ResponseDto{
			Status:     "404",
			Error:      errorMessage,
			Message:    h.Source.GetMessage(errorMessage),
			StackTrace: er.Error(),
		},
	)

	return
}
