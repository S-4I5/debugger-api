package redis

import (
	"context"
	def "debugger-api/internal/repository"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ def.Repository = (*postgresRepository)(nil)

const tableName = "data"

type postgresRepository struct {
	client *pgxpool.Pool
}

func NewPostgresRepository(client *pgxpool.Pool) *postgresRepository {
	return &postgresRepository{
		client: client,
	}
}

func (r *postgresRepository) Create(cxt context.Context, data string) (uuid.UUID, error) {
	const op = "repository/postgres/create"

	builder := squirrel.Insert(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns("content").
		Values(data).
		Suffix("RETURNING id")

	sql, args, err := builder.ToSql()
	if err != nil {
		return uuid.Nil, fmt.Errorf(op+" : error while building sql : %s", err.Error())
	}

	rows, err := r.client.Query(cxt, sql, args...)
	if err != nil {
		return uuid.Nil, fmt.Errorf(op+" : error while exec query : %s", err.Error())
	}
	defer rows.Close()

	rows.Next()
	var result uuid.UUID
	_ = rows.Scan(&result)

	err = rows.Err()
	if err != nil {
		return uuid.Nil, fmt.Errorf(op + " : error after exec query sql")
	}

	return result, nil
}

func (r *postgresRepository) Delete(cxt context.Context, id uuid.UUID) error {
	const op = "repository/postgres/delete"

	builder := squirrel.Delete(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": id}).
		Suffix("RETURNING id")

	sql, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf(op+" : error while building sql : %s", err.Error())
	}

	rows, err := r.client.Query(cxt, sql, args...)
	if err != nil {
		return fmt.Errorf(op+" : error while exec query sql : %s", err.Error())
	}
	defer rows.Close()

	if !rows.Next() && rows.Err() == nil {
		return fmt.Errorf(op+": there is no data with id : %s", id)
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresRepository) Update(cxt context.Context, data string, id uuid.UUID) error {
	const op = "repository/postgres/update"

	builder := squirrel.Update(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Set("content", data).
		Where(squirrel.Eq{"id": id}).
		Suffix("RETURNING id")

	sql, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf(op+" : error while building sql : %s", err.Error())
	}

	rows, err := r.client.Query(cxt, sql, args...)
	if err != nil {
		return fmt.Errorf(op+" : error while exec query sql : %s", err.Error())
	}
	defer rows.Close()

	if !rows.Next() && rows.Err() == nil {
		return fmt.Errorf(op+" : there is no data with id : %s", id)
	}

	err = rows.Err()
	if err != nil {
		return fmt.Errorf(op + " : error after exec query sql")
	}

	return nil
}

func (r *postgresRepository) Get(cxt context.Context, id uuid.UUID) (string, error) {
	const op = "repository/postgres/get"

	builder := squirrel.Select("content").
		PlaceholderFormat(squirrel.Dollar).
		From(tableName).
		Where(squirrel.Eq{"id": id})

	sql, args, err := builder.ToSql()
	if err != nil {
		return "", fmt.Errorf(op+" : error while building sql : %s", err.Error())
	}

	rows, err := r.client.Query(cxt, sql, args...)
	if err != nil {
		return "", fmt.Errorf(op+" : error while exec query sql : %s", err.Error())
	}
	defer rows.Close()

	rows.Next()
	var result string

	err = rows.Scan(&result)
	if err != nil {
		return "", fmt.Errorf(op+" : there is no data with id : %s", id)
	}

	err = rows.Err()
	if err != nil {
		return "", fmt.Errorf(op + " : error after exec query sql")
	}

	return result, nil
}
