package data

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *Service) Get(ctx context.Context, key string) (map[string]string, error) {
	const op = "service/get"

	data, err := s.dataRepository.Get(ctx, key)
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
