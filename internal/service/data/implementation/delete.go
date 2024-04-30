package implementation

import (
	"context"
	"fmt"
)

func (s *Service) Delete(ctx context.Context, key string) error {
	const op = "service/delete"

	err := s.dataRepository.Delete(ctx, key)
	if err != nil {
		return fmt.Errorf(op+": Cannot delete data: %v\n", err)
	}

	return nil
}
