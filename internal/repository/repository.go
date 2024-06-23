package repository

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	Create(_ context.Context, data string) (uuid.UUID, error)
	Delete(_ context.Context, id uuid.UUID) error
	Update(_ context.Context, data string, id uuid.UUID) error
	Get(_ context.Context, id uuid.UUID) (string, error)
}
