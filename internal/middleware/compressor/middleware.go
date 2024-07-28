package compressor

import (
	"net/http"
)

type provider struct {
}

func NewMiddlewareProvider() *provider {
	return &provider{}
}

// Doesn't work for now :(
func (p *provider) GetCompressorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		next.ServeHTTP(writer, request)

		//writer.Header().Set("Content-Encoding", "br")
	})
}
