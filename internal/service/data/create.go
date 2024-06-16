package data

import (
	"context"
	"fmt"
)

func (s *Service) Create(ctx context.Context, data string, key string) error {
	const op = "service/create"

	err := s.dataRepository.Create(ctx, data, key)
	if err != nil {
		return fmt.Errorf(op+": Cannot create data: %v\n", err)
	}

	return nil
}
