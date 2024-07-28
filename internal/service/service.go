package service

import (
	"context"
	"debugger-api/internal/model"
	"debugger-api/internal/model/dto"
	"github.com/google/uuid"
)

type MockService interface {
	Create(ctx context.Context, dto dto.CreateMockDto) (dto.MockDto, error)
	Get(ctx context.Context, id uuid.UUID) (dto.MockDto, error)
	Update(ctx context.Context, dto dto.UpdateMockDto, id uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetContent(ctx context.Context, id uuid.UUID) (model.Json, error)
}
