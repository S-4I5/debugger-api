package requestid

import (
	"context"
	"debugger-api/internal/model"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

const requestIdKey = "requestId"

var (
	errParseRequestIdFromContext = fmt.Errorf("cannot extract id from context")
)

type provider struct {
}

func NewMiddlewareProvider() *provider {
	return &provider{}
}

func (p *provider) GetRequestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requestId, _ := uuid.NewRandom()

		ctxWithId := context.WithValue(request.Context(), requestIdKey, requestId)

		next.ServeHTTP(writer, request.WithContext(ctxWithId))
	})
}

func GetRequestIdFromContext(ctx context.Context) (uuid.UUID, error) {
	const op = "middleware/request_id/get"
	id, ok := ctx.Value(requestIdKey).(uuid.UUID)
	if !ok {
		return uuid.Nil, model.BuildSubErrorWithOperation(op, errParseRequestIdFromContext.Error())
	}
	return id, nil
}
