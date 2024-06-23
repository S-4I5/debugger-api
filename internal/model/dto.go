package model

import "github.com/google/uuid"

type PostResponse struct {
	Id uuid.UUID `json:"id"`
}
