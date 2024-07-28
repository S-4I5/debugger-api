package fullurl

import (
	"context"
	"debugger-api/internal/model"
	"fmt"
	"net/http"
)

const fullUrlKey = "fullUrl"

var (
	errParseFullValueFromContext = fmt.Errorf("cannot extract full path from context")
)

type provider struct {
}

func NewMiddlewareProvider() *provider {
	return &provider{}
}

func (p *provider) GetFullRequestUrlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		newCtx := context.WithValue(request.Context(), fullUrlKey, request.URL.Path)

		next.ServeHTTP(writer, request.WithContext(newCtx))
	})
}

func GetFullRequestUrlFromContext(ctx context.Context) (string, error) {
	const op = "middleware/full_url/get"
	path, ok := ctx.Value(fullUrlKey).(string)
	if !ok {
		return "", model.BuildSubErrorWithOperation(op, errParseFullValueFromContext.Error())
	}
	return path, nil
}
