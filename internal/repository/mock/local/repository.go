package local

import (
	"context"
	"debugger-api/internal/httperr"
	model2 "debugger-api/internal/model"
	"debugger-api/internal/model/entity"
	def "debugger-api/internal/repository"
	"debugger-api/internal/repository/mock/local/mapper"
	"debugger-api/internal/repository/mock/local/model"
	"fmt"
	"github.com/google/uuid"
	"sync"
	"time"
)

var _ def.MockRepository = (*repository)(nil)

type repository struct {
	data map[uuid.UUID]local_model.Mock
	m    sync.RWMutex
}

var noMockFoundByGivenId = fmt.Errorf("no mock found by given id")

const errorNoMockFound = "error while searching for mock"

func NewRepository() *repository {
	return &repository{
		data: make(map[uuid.UUID]local_model.Mock),
	}
}

func (r *repository) Save(_ context.Context, mock entity.Mock) (entity.Mock, error) {
	const op = "repository/local/create"

	r.m.Lock()
	defer r.m.Unlock()

	id, _ := uuid.NewRandom()

	_, isPresent := r.data[id]
	for isPresent {
		id, _ = uuid.NewRandom()
		_, isPresent = r.data[id]
	}

	dbMock := local_mapper.MockToDbMock(mock)

	r.data[id] = dbMock

	return local_mapper.DbMockToMock(dbMock, id), nil
}

func (r *repository) Delete(_ context.Context, id uuid.UUID) error {
	const op = "repository/local/delete"

	r.m.Lock()
	defer r.m.Unlock()

	_, isPresent := r.data[id]
	if !isPresent {
		return model2.NewServiceError(
			model2.BuildErrorWithOperation(op, errorNoMockFound, noMockFoundByGivenId),
			httperr.MockNotFoundError,
		)
	}

	delete(r.data, id)

	return nil
}

func (r *repository) UpdateContent(_ context.Context, json model2.Json, id uuid.UUID) error {
	const op = "repository/local/update"

	r.m.Lock()
	defer r.m.Unlock()

	mock, isPresent := r.data[id]
	if !isPresent {
		return model2.NewServiceError(
			model2.BuildErrorWithOperation(op, errorNoMockFound, noMockFoundByGivenId),
			httperr.MockNotFoundError,
		)
	}

	now := time.Now()

	mock.Content = json
	mock.UpdatedAt = &now
	r.data[id] = mock

	return nil
}

func (r *repository) Get(_ context.Context, id uuid.UUID) (entity.Mock, error) {
	const op = "repository/local/get"

	r.m.RLock()
	defer r.m.RUnlock()

	mock, ok := r.data[id]
	if !ok {
		return entity.Mock{}, model2.NewServiceError(
			model2.BuildErrorWithOperation(op, errorNoMockFound, noMockFoundByGivenId),
			httperr.MockNotFoundError,
		)
	}

	return local_mapper.DbMockToMock(mock, id), nil
}
