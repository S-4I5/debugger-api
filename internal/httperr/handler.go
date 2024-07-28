package httperr

import (
	"debugger-api/internal/middleware/fullurl"
	"debugger-api/internal/middleware/requestid"
	"debugger-api/internal/model"
	"github.com/go-chi/render"
	"net/http"
	"time"
)

type Handler interface {
	HandleServiceError(err error, w http.ResponseWriter, r *http.Request)
	HandleUnprocessableEntityError(err error, w http.ResponseWriter, r *http.Request)
	HandleIncorrectRequestParamError(err error, w http.ResponseWriter, r *http.Request)
	HandleResponseProcessingError(err error, w http.ResponseWriter, r *http.Request)
}

type handler struct {
	Source
}

func NewErrorResponseHandler(source Source) *handler {
	return &handler{source}
}

func (h *handler) HandleServiceError(err error, w http.ResponseWriter, r *http.Request) {
	serviceError, ok := err.(*model.ServiceError)
	if !ok {
		h.handleError(model.NewServiceError(err, InternalServiceError), w, r, 500)
	}

	h.handleError(serviceError, w, r, 400)
}

func (h *handler) HandleResponseProcessingError(err error, w http.ResponseWriter, r *http.Request) {
	h.handleError(model.NewServiceError(err, ResponseProcessingError), w, r, 500)
}

func (h *handler) HandleUnprocessableEntityError(err error, w http.ResponseWriter, r *http.Request) {
	h.handleError(model.NewServiceError(err, UnprocessableEntityError), w, r, 420)
}

func (h *handler) HandleIncorrectRequestParamError(err error, w http.ResponseWriter, r *http.Request) {
	h.handleError(model.NewServiceError(err, IncorrectRequestParamError), w, r, 400)
}

func (h *handler) handleError(err *model.ServiceError, w http.ResponseWriter, r *http.Request, status int) {
	requestId, _ := requestid.GetRequestIdFromContext(r.Context())

	fullPath, _ := fullurl.GetFullRequestUrlFromContext(r.Context())

	render.Status(r, status)
	render.JSON(w, r,
		ResponseDto{
			Status:      status,
			RequestId:   requestId.String(),
			Message:     h.Source.GetMessage(err.Message),
			MessageCode: err.Message,
			StackTrace:  err.ErrorStringSlice(),
			Timestamp:   time.Now(),
			Path:        fullPath,
		},
	)
}
