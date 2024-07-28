package dto

import (
	"debugger-api/internal/model"
	"github.com/google/uuid"
	"time"
)

type Json map[string]interface{}

type CreateMockDto struct {
	Content model.Json `json:"content" validate:"required"`
}

type UpdateMockDto struct {
	NewContent model.Json `json:"newContent" validate:"required"`
}

type MockDto struct {
	Id        uuid.UUID  `json:"id"`
	Content   model.Json `json:"content"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
