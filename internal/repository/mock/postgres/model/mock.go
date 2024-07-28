package postgres_model

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Mock struct {
	Id        uuid.UUID
	Content   string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
