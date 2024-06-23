package local

import (
	"context"
	def "debugger-api/internal/repository"
	"fmt"
	"github.com/google/uuid"
	"sync"
)

var _ def.Repository = (*localRepository)(nil)

type localRepository struct {
	data map[uuid.UUID]string
	m    sync.RWMutex
}

func NewLocalRepository() *localRepository {
	return &localRepository{
		data: make(map[uuid.UUID]string),
	}
}

func (r *localRepository) Create(_ context.Context, data string) (uuid.UUID, error) {
	const op = "repository/local/create"

	r.m.Lock()
	defer r.m.Unlock()

	id, _ := uuid.NewRandom()

	_, isPresent := r.data[id]
	for isPresent {
		id, _ = uuid.NewRandom()
		_, isPresent = r.data[id]
	}

	r.data[id] = data

	return id, nil
}

func (r *localRepository) Delete(_ context.Context, id uuid.UUID) error {
	const op = "repository/local/delete"

	r.m.Lock()
	defer r.m.Unlock()

	_, isPresent := r.data[id]
	if !isPresent {
		return fmt.Errorf(op + ": value under give key does not exists")
	}

	delete(r.data, id)

	return nil
}

func (r *localRepository) Update(_ context.Context, data string, id uuid.UUID) error {
	const op = "repository/local/update"

	r.m.Lock()
	defer r.m.Unlock()

	_, isPresent := r.data[id]
	if !isPresent {
		return fmt.Errorf(op + ": value under give key does not exists")
	}

	r.data[id] = data

	return nil
}

func (r *localRepository) Get(_ context.Context, id uuid.UUID) (string, error) {
	const op = "repository/local/get"

	r.m.RLock()
	defer r.m.RUnlock()

	data, ok := r.data[id]
	if !ok {
		return "", fmt.Errorf(op + ": value under give key does not exists")
	}

	return data, nil
}
