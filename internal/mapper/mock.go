package mapper

import (
	"debugger-api/internal/model/dto"
	"debugger-api/internal/model/entity"
	"time"
)

func MockToMockDto(mock entity.Mock) dto.MockDto {
	return dto.MockDto{
		Id:        mock.Id,
		Content:   mock.Content,
		CreatedAt: mock.CreatedAt,
		UpdatedAt: mock.UpdatedAt,
	}
}

func CreateMockDtoToMock(mockDto dto.CreateMockDto) entity.Mock {
	return entity.Mock{
		Content:   mockDto.Content,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}
}

func MockDtoToMock(mock dto.MockDto) entity.Mock {
	return entity.Mock{
		Id:        mock.Id,
		Content:   mock.Content,
		CreatedAt: mock.CreatedAt,
		UpdatedAt: mock.UpdatedAt,
	}
}
