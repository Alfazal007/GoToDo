// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user_queries.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
insert into users (id, name, age, password) values ($1, $2, $3, $4) returning id, name, age, password
`

type CreateUserParams struct {
	ID       uuid.UUID
	Name     string
	Age      int32
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Age,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Age,
		&i.Password,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
select id, name, age, password from users where id=$1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Age,
		&i.Password,
	)
	return i, err
}

const getUserByName = `-- name: GetUserByName :one
select id, name, age, password from users where name=$1
`

func (q *Queries) GetUserByName(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByName, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Age,
		&i.Password,
	)
	return i, err
}
