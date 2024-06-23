package data

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

func (s *service) Update(ctx context.Context, data string, id uuid.UUID) error {
	const op = "service/update"

	err := s.dataRepository.Update(ctx, data, id)
	if err != nil {
		return fmt.Errorf(op+": Cannot update data: %v\n", err)
	}

	return nil
}
