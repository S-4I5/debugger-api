package repository

import (
	"context"
	"debugger-api/internal/model"
	"debugger-api/internal/model/entity"
	"github.com/google/uuid"
)

type MockRepository interface {
	Save(ctx context.Context, mock entity.Mock) (entity.Mock, error)
	Delete(ctx context.Context, id uuid.UUID) error
	UpdateContent(ctx context.Context, json model.Json, id uuid.UUID) error
	Get(ctx context.Context, id uuid.UUID) (entity.Mock, error)
}
