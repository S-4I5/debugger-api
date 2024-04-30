package error

type ResponseDto struct {
	Status     string `json:"status"`
	Message    string `json:"message,omitempty"`
	Error      string `json:"error,omitempty"`
	StackTrace string `json:"stackTrace,omitempty"`
}
