package postgres

import (
	"context"
	"debugger-api/internal/httperr"
	"debugger-api/internal/model"
	"debugger-api/internal/model/entity"
	def "debugger-api/internal/repository"
	postgres_mapper "debugger-api/internal/repository/mock/postgres/mapper"
	postgres_model "debugger-api/internal/repository/mock/postgres/model"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

var _ def.MockRepository = (*repository)(nil)

type repository struct {
	client *pgxpool.Pool
}

const (
	tableName           = "mock"
	columnCreatedAtName = "created_at"
	columnUpdatedAtName = "updated_at"
	columnContentName   = "content"
	columnIdName        = "id"
)

var noMockFoundByGivenId = fmt.Errorf("no mock found by given id")

const (
	errorWhileExecutingQuery  = "error while executing query"
	errorWhileBuildingQuery   = "error while building query"
	errorWhileReadingResponse = "error while reading response"
	errorNoMockFound          = "error while searching for mock"
)

func NewRepository(client *pgxpool.Pool) *repository {
	return &repository{
		client: client,
	}
}

func (r *repository) Save(_ context.Context, mock entity.Mock) (entity.Mock, error) {
	const op = "repository/postgres/create"

	ctx := context.TODO()

	builder := squirrel.Insert(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(columnContentName, columnCreatedAtName).
		Values(mock.Content, mock.CreatedAt).
		Suffix("RETURNING id")

	sql, args, err := builder.ToSql()
	if err != nil {
		return entity.Mock{}, model.BuildErrorWithOperation(op, errorWhileBuildingQuery, err)
	}

	rows, err := r.client.Query(ctx, sql, args...)
	if err != nil {
		return entity.Mock{},
			model.BuildErrorWithOperation(op, errorWhileExecutingQuery, err)
	}
	defer rows.Close()

	rows.Next()

	var id uuid.UUID
	err = rows.Scan(&id)
	if err != nil {
		return entity.Mock{}, model.BuildErrorWithOperation(op, errorWhileReadingResponse, err)
	}

	mock.Id = id

	return mock, rows.Err()
}

func (r *repository) Delete(_ context.Context, id uuid.UUID) error {
	const op = "repository/postgres/delete"

	ctx := context.TODO()

	builder := squirrel.Delete(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{columnIdName: id})

	sql, args, err := builder.ToSql()
	if err != nil {
		return model.BuildErrorWithOperation(op, errorWhileBuildingQuery, err)
	}

	result, err := r.client.Exec(ctx, sql, args...)
	if err != nil {
		return model.BuildErrorWithOperation(op, errorWhileExecutingQuery, err)
	}

	if result.RowsAffected() == 0 {
		return model.NewServiceError(
			model.BuildErrorWithOperation(op, errorNoMockFound, noMockFoundByGivenId),
			httperr.MockNotFoundError,
		)
	}

	return nil
}

func (r *repository) UpdateContent(_ context.Context, js model.Json, id uuid.UUID) error {
	const op = "repository/postgres/update"

	ctx := context.TODO()

	stringJson, err := json.Marshal(js)
	if err != nil {
		return err
	}

	builder := squirrel.Update(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Set(columnContentName, stringJson).
		Set(columnUpdatedAtName, time.Now()).
		Where(squirrel.Eq{columnIdName: id})

	sql, args, err := builder.ToSql()
	if err != nil {
		return model.BuildErrorWithOperation(op, errorWhileBuildingQuery, err)
	}

	result, err := r.client.Exec(ctx, sql, args...)
	if err != nil {
		return model.BuildErrorWithOperation(op, errorWhileExecutingQuery, err)
	}

	if result.RowsAffected() == 0 {
		return model.NewServiceError(
			model.BuildErrorWithOperation(op, errorNoMockFound, noMockFoundByGivenId),
			httperr.MockNotFoundError,
		)
	}

	return nil
}

func (r *repository) Get(_ context.Context, id uuid.UUID) (entity.Mock, error) {
	const op = "repository/postgres/get"

	ctx := context.TODO()

	builder := squirrel.Select("*").
		PlaceholderFormat(squirrel.Dollar).
		From(tableName).
		Where(squirrel.Eq{columnIdName: id})

	sql, args, err := builder.ToSql()
	if err != nil {
		return entity.Mock{},
			model.BuildErrorWithOperation(op, errorWhileBuildingQuery, err)
	}

	rows, err := r.client.Query(ctx, sql, args...)
	if err != nil {
		return entity.Mock{},
			model.BuildErrorWithOperation(op, errorWhileExecutingQuery, err)
	}
	defer rows.Close()

	mock, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[postgres_model.Mock])
	if err != nil {
		return entity.Mock{}, model.NewServiceError(
			model.BuildErrorWithOperation(op, errorNoMockFound, err),
			httperr.MockNotFoundError,
		)
	}

	return postgres_mapper.DbMockToMock(mock)
}
