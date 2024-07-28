package mock

import (
	"context"
	"debugger-api/internal/mapper"
	"debugger-api/internal/model"
	"debugger-api/internal/model/dto"
	"debugger-api/internal/repository"
	def "debugger-api/internal/service"
	"github.com/google/uuid"
)

var _ def.MockService = (*service)(nil)

type service struct {
	mockRepository repository.MockRepository
}

const (
	createMockError = "error while creating mock"
	deleteMockError = "error while deleting mock"
	getMockError    = "get mock error"
	updateMockError = "error while updating mock"
)

func NewDataService(dataRepository repository.MockRepository) *service {
	return &service{
		mockRepository: dataRepository,
	}
}

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	const op = "service/delete"

	err := s.mockRepository.Delete(ctx, id)
	if err != nil {
		return model.FromError(err, model.BuildSubErrorWithOperation(op, deleteMockError))
	}

	return nil
}

func (s *service) Get(ctx context.Context, id uuid.UUID) (dto.MockDto, error) {
	const op = "service/get"

	mock, err := s.mockRepository.Get(ctx, id)
	if err != nil {
		return dto.MockDto{}, model.FromError(err, model.BuildSubErrorWithOperation(op, getMockError))
	}

	return mapper.MockToMockDto(mock), nil
}

func (s *service) Update(ctx context.Context, updateDto dto.UpdateMockDto, id uuid.UUID) error {
	const op = "service/update"

	err := s.mockRepository.UpdateContent(ctx, updateDto.NewContent, id)
	if err != nil {
		return model.FromError(err, model.BuildSubErrorWithOperation(op, updateMockError))
	}

	return nil
}

func (s *service) Create(ctx context.Context, createDto dto.CreateMockDto) (dto.MockDto, error) {
	const op = "service/create"

	mock := mapper.CreateMockDtoToMock(createDto)

	saved, err := s.mockRepository.Save(ctx, mock)
	if err != nil {
		return dto.MockDto{}, model.FromError(err, model.BuildSubErrorWithOperation(op, createMockError))
	}

	return mapper.MockToMockDto(saved), nil
}

func (s *service) GetContent(ctx context.Context, id uuid.UUID) (model.Json, error) {
	const op = "service/get_context"

	mock, err := s.mockRepository.Get(ctx, id)
	if err != nil {
		return nil, model.FromError(err, model.BuildSubErrorWithOperation(op, getMockError))
	}

	return mock.Content, nil
}
