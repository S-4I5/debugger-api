package httperr

import "time"

const (
	UnprocessableEntityError   = "api.request.body.unprocessable"
	IncorrectRequestParamError = "api.request.param.incorrect"
	InternalServiceError       = "api.internal"
	ResponseProcessingError    = "api.response.cannot_process"
	MockNotFoundError          = "api.mock.not_found"
)

type ResponseDto struct {
	Status      int       `json:"status"`
	Message     string    `json:"message,omitempty"`
	MessageCode string    `json:"messageCode"`
	StackTrace  []string  `json:"stackTrace,omitempty"`
	RequestId   string    `json:"requestId"`
	Timestamp   time.Time `json:"timestamp"`
	Path        string    `json:"path"`
}
