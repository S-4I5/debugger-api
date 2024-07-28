package redis

import (
	"context"
	"debugger-api/internal/httperr"
	"debugger-api/internal/model"
	"debugger-api/internal/model/entity"
	def "debugger-api/internal/repository"
	"debugger-api/internal/repository/mock/redis/mapper"
	model2 "debugger-api/internal/repository/mock/redis/model"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"time"
)

var _ def.MockRepository = (*repository)(nil)

type repository struct {
	client *redis.Client
}

var noMockFoundByGivenId = fmt.Errorf("no mock found by given id")

const errorDelete = "error while trying to delete mock"

const (
	errorSetMock     = "cannot save mock"
	errorParseMock   = "error while parsing mock"
	errorNoMockFound = "error while searching for mock"
)

func NewRepository(client *redis.Client) *repository {
	return &repository{
		client: client,
	}
}

func (r *repository) Save(_ context.Context, mock entity.Mock) (entity.Mock, error) {
	const op = "repository/redis/create"

	ctx := context.TODO()

	id, _ := uuid.NewRandom()

	result, err := r.client.Exists(ctx, id.String()).Result()
	for err != nil || result == 1 {
		if err != nil {
			return entity.Mock{}, model.BuildErrorWithOperation(op, errorSetMock, err)
		}
		id, _ = uuid.NewRandom()
		result, err = r.client.Exists(ctx, id.String()).Result()
	}

	err = r.client.Set(ctx, id.String(), redis_mapper.MockToDbMock(mock), 0).Err()
	if err != nil {
		return entity.Mock{}, model.BuildErrorWithOperation(op, errorSetMock, err)
	}

	mock.Id = id

	return mock, nil
}

func (r *repository) Delete(_ context.Context, id uuid.UUID) error {
	const op = "repository/redis/delete"

	ctx := context.TODO()

	result, err := r.client.Del(ctx, id.String()).Result()
	if err != nil || result == 0 {
		return model.BuildErrorWithOperation(op, errorDelete, err)
	} else if result == 0 {
		return model.NewServiceError(
			model.BuildErrorWithOperation(op, errorNoMockFound, noMockFoundByGivenId),
			httperr.MockNotFoundError,
		)
	}

	return nil
}

func (r *repository) UpdateContent(_ context.Context, js model.Json, id uuid.UUID) error {
	const op = "repository/redis/update"

	ctx := context.TODO()

	rowMock, err := r.client.Get(ctx, id.String()).Result()
	if err != nil {
		return model.NewServiceError(
			model.BuildErrorWithOperation(op, errorNoMockFound, noMockFoundByGivenId),
			httperr.MockNotFoundError,
		)
	}

	var mock model2.Mock

	err = json.Unmarshal([]byte(rowMock), &mock)
	if err != nil {
		return model.BuildErrorWithOperation(op, errorParseMock, err)
	}

	now := time.Now()

	mock.Content = js
	mock.UpdatedAt = &now

	err = r.client.Set(ctx, id.String(), mock, 0).Err()
	if err != nil {
		return model.BuildErrorWithOperation(op, errorSetMock, err)
	}

	return nil
}

func (r *repository) Get(_ context.Context, id uuid.UUID) (entity.Mock, error) {
	const op = "repository/redis/get"

	ctx := context.TODO()

	rowMock, err := r.client.Get(ctx, id.String()).Result()
	if err != nil {
		return entity.Mock{}, model.NewServiceError(
			model.BuildErrorWithOperation(op, errorNoMockFound, noMockFoundByGivenId),
			httperr.MockNotFoundError,
		)
	}

	var mock model2.Mock

	err = json.Unmarshal([]byte(rowMock), &mock)
	if err != nil {
		return entity.Mock{}, model.BuildErrorWithOperation(op, errorParseMock, err)
	}

	return redis_mapper.DbMockToMock(mock, id), nil
}
