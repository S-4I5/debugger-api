package implementation

import (
	"context"
	"fmt"
)

func (s *Service) Update(ctx context.Context, data string, key string) error {
	const op = "service/update"

	err := s.dataRepository.Update(ctx, data, key)
	if err != nil {
		return fmt.Errorf(op+": Cannot update data: %v\n", err)
	}

	return nil
}
