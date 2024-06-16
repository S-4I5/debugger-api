package data

import (
	"context"
	def "debugger-api/internal/repository"
	"fmt"
	"sync"
)

var _ def.Repository = (*Repository)(nil)

type Repository struct {
	data map[string]string
	m    sync.RWMutex
}

func NewRepository() *Repository {
	return &Repository{
		data: make(map[string]string),
	}
}

func (r *Repository) Create(_ context.Context, data string, key string) error {
	const op = "repository/create"

	r.m.Lock()
	defer r.m.Unlock()

	_, isPresent := r.data[key]
	if isPresent {
		return fmt.Errorf(op + ": value under give key already exists")
	}

	r.data[key] = data

	return nil
}

func (r *Repository) Delete(_ context.Context, key string) error {
	const op = "repository/delete"

	r.m.Lock()
	defer r.m.Unlock()

	_, isPresent := r.data[key]
	if !isPresent {
		return fmt.Errorf(op + ": value under give key does not exists")
	}

	delete(r.data, key)

	return nil
}

func (r *Repository) Update(_ context.Context, data string, key string) error {
	const op = "repository/update"

	r.m.Lock()
	defer r.m.Unlock()

	_, isPresent := r.data[key]
	if !isPresent {
		return fmt.Errorf(op + ": value under give key does not exists")
	}

	r.data[key] = data

	return nil
}

func (r *Repository) Get(_ context.Context, key string) (string, error) {
	const op = "repository/get"

	r.m.RLock()
	defer r.m.RUnlock()

	data, ok := r.data[key]
	if !ok {
		return "", fmt.Errorf(op + ": value under give key does not exists")
	}

	return data, nil
}
