package redis_mapper

import (
	"debugger-api/internal/model/entity"
	"debugger-api/internal/repository/mock/redis/model"
	"github.com/google/uuid"
)

func MockToDbMock(mock entity.Mock) redis_model.Mock {
	return redis_model.Mock{
		Content:   mock.Content,
		CreatedAt: mock.CreatedAt,
		UpdatedAt: mock.UpdatedAt,
	}
}

func DbMockToMock(mock redis_model.Mock, id uuid.UUID) entity.Mock {
	return entity.Mock{
		Id:        id,
		Content:   mock.Content,
		CreatedAt: mock.CreatedAt,
		UpdatedAt: mock.UpdatedAt,
	}
}
