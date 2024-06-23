package data

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

func (s *service) Create(ctx context.Context, data string) (uuid.UUID, error) {
	const op = "service/create"

	id, err := s.dataRepository.Create(ctx, data)
	if err != nil {
		return uuid.Nil, fmt.Errorf(op+": Cannot create data: %v\n", err)
	}

	return id, nil
}
