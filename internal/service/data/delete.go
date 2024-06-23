package data

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	const op = "service/delete"

	err := s.dataRepository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf(op+": Cannot delete data: %v\n", err)
	}

	return nil
}
