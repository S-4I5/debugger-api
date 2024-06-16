package controller

import (
	"context"
	"net/http"
)

type Controller interface {
	PostData(cxt context.Context) http.HandlerFunc
	GetData(ctx context.Context) http.HandlerFunc
	UpdateData(cxt context.Context) http.HandlerFunc
	DeleteData(ctx context.Context) http.HandlerFunc
}
