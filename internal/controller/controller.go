package controller

import (
	"context"
	"net/http"
)

type MockController interface {
	PostMock(cxt context.Context) http.HandlerFunc
	GetMock(ctx context.Context) http.HandlerFunc
	UpdateMock(cxt context.Context) http.HandlerFunc
	DeleteMock(ctx context.Context) http.HandlerFunc
	GetMockContent(ctx context.Context) http.HandlerFunc
}
