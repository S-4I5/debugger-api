package redis

import (
	"context"
	def "debugger-api/internal/repository"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var _ def.Repository = (*redisRepository)(nil)

type redisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *redisRepository {
	return &redisRepository{
		client: client,
	}
}

func (r *redisRepository) Create(cxt context.Context, data string) (uuid.UUID, error) {
	const op = "repository/redis/create"

	id, _ := uuid.NewRandom()

	result, err := r.client.Exists(cxt, id.String()).Result()
	for err != nil || result == 1 {
		if err != nil {
			return uuid.Nil, fmt.Errorf(op+": error while searching for value under key %s : %s", id, err.Error())
		}
		id, _ = uuid.NewRandom()
		result, err = r.client.Exists(cxt, id.String()).Result()
	}

	err = r.client.Set(cxt, id.String(), data, 0).Err()
	if err != nil {
		return uuid.Nil, fmt.Errorf(op + ": value under give key already exists")
	}

	return id, nil
}

func (r *redisRepository) Delete(cxt context.Context, id uuid.UUID) error {
	const op = "repository/redis/delete"

	result, err := r.client.Del(cxt, id.String()).Result()
	if err != nil {
		return fmt.Errorf(op + ": value under give key does not exists")
	}
	if result == 0 {
		return fmt.Errorf(op+": no value under given key %s", id)
	}

	return nil
}

func (r *redisRepository) Update(cxt context.Context, data string, id uuid.UUID) error {
	const op = "repository/redis/update"

	result, err := r.client.Exists(cxt, id.String()).Result()
	if err != nil {
		return fmt.Errorf(op+": error while searching for value under key %s : %s", id, err.Error())
	}
	if result == 0 {
		return fmt.Errorf(op+": no value under given key %s", id)
	}

	err = r.client.Set(cxt, id.String(), data, 0).Err()
	if err != nil {
		return fmt.Errorf(op + ": value under give key already exists")
	}

	return nil
}

func (r *redisRepository) Get(cxt context.Context, id uuid.UUID) (string, error) {
	const op = "repository/redis/get"

	result, err := r.client.Get(cxt, id.String()).Result()
	if err != nil {
		return "", fmt.Errorf(op + ": value under give key does not exists")
	}

	return result, nil
}
