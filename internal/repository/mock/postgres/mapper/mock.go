package postgres_mapper

import (
	"database/sql"
	"debugger-api/internal/model"
	"debugger-api/internal/model/entity"
	"debugger-api/internal/repository/mock/postgres/model"
	"encoding/json"
)

func MockToDbMock(mock entity.Mock) (postgres_model.Mock, error) {
	stringContent, err := json.Marshal(mock.Content)
	if err != nil {
		return postgres_model.Mock{}, err
	}

	updatedAt := sql.NullTime{
		Time:  *mock.UpdatedAt,
		Valid: mock.UpdatedAt == nil,
	}

	return postgres_model.Mock{
		Content:   string(stringContent),
		CreatedAt: mock.CreatedAt,
		UpdatedAt: updatedAt,
	}, nil
}

func DbMockToMock(mock postgres_model.Mock) (entity.Mock, error) {
	var mapContent model.Json
	err := json.Unmarshal([]byte(mock.Content), &mapContent)
	if err != nil {
		return entity.Mock{}, err
	}

	return entity.Mock{
		Id:        mock.Id,
		Content:   mapContent,
		CreatedAt: mock.CreatedAt,
		UpdatedAt: &mock.UpdatedAt.Time,
	}, nil
}
