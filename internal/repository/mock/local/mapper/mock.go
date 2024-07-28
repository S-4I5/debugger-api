package local_mapper

import (
	"debugger-api/internal/model/entity"
	"debugger-api/internal/repository/mock/local/model"
	"github.com/google/uuid"
)

func MockToDbMock(mock entity.Mock) local_model.Mock {
	return local_model.Mock{
		Content:   mock.Content,
		CreatedAt: mock.CreatedAt,
		UpdatedAt: mock.UpdatedAt,
	}
}

func DbMockToMock(mock local_model.Mock, id uuid.UUID) entity.Mock {
	return entity.Mock{
		Id:        id,
		Content:   mock.Content,
		CreatedAt: mock.CreatedAt,
		UpdatedAt: mock.UpdatedAt,
	}
}
