package entity

import (
	"debugger-api/internal/model"
	"github.com/google/uuid"
	"time"
)

type Mock struct {
	Id        uuid.UUID
	Content   model.Json
	CreatedAt time.Time
	UpdatedAt *time.Time
}
