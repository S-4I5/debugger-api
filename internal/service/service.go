package service

import (
	"context"
)

type Service interface {
	Create(ctx context.Context, data string, key string) error
	Get(ctx context.Context, key string) (map[string]string, error)
	Update(ctx context.Context, data string, key string) error
	Delete(ctx context.Context, key string) error
}
