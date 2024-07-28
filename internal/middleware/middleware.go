package middleware

import "net/http"

type Middleware func(next http.Handler) http.Handler

type FullRequestUrlMiddlewareProvider interface {
	GetFullRequestUrlMiddleware(next http.Handler) http.Handler
}

type RequestIdMiddlewareProvider interface {
	GetRequestIdMiddleware(next http.Handler) http.Handler
}

type CompressorMiddlewareProvider interface {
	GetCompressorMiddleware(next http.Handler) http.Handler
}

func CreateStack(stack ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(stack) - 1; i >= 0; i-- {
			cur := stack[i]
			next = cur(next)
		}

		return next
	}
}
