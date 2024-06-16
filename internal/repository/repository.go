package repository

import (
	"context"
)

type Repository interface {
	Create(_ context.Context, data string, key string) error
	Delete(_ context.Context, key string) error
	Update(_ context.Context, data string, key string) error
	Get(_ context.Context, key string) (string, error)
}
