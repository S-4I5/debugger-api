package local_model

import (
	"debugger-api/internal/model"
	"time"
)

type Mock struct {
	Content   model.Json
	CreatedAt time.Time
	UpdatedAt *time.Time
}
