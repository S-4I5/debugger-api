package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

func (s *service) Get(ctx context.Context, id uuid.UUID) (map[string]string, error) {
	const op = "service/get"

	data, err := s.dataRepository.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf(op+": Cannot get data: %v\n", err)
	}

	var dataJsonMap map[string]string
	err = json.Unmarshal([]byte(data), &dataJsonMap)
	if err != nil {
		return nil, fmt.Errorf(op+": Cannot convert data: %v\n", err)
	}

	return dataJsonMap, nil
}
