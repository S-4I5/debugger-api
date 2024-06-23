package service

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, data string) (uuid.UUID, error)
	Get(ctx context.Context, id uuid.UUID) (map[string]string, error)
	Update(ctx context.Context, data string, id uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID) error
}
