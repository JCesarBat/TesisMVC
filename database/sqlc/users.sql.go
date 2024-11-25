// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
	"database/sql"
)

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "users" WHERE id =$1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, id_municipio, username, password, email, super_user, created_at FROM users WHERE "username" = $1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.IDMunicipio,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.SuperUser,
		&i.CreatedAt,
	)
	return i, err
}

const getUserID = `-- name: GetUserID :one
SELECT id, id_municipio, username, password, email, super_user, created_at FROM users WHERE "id" = $1
`

func (q *Queries) GetUserID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.IDMunicipio,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.SuperUser,
		&i.CreatedAt,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO "users"(
                    "username",
                    "password",
                    "email",
                    "super_user",
                    "id_municipio"
) VALUES ($1, $2, $3, $4,$5) RETURNING id, id_municipio, username, password, email, super_user, created_at
`

type InsertUserParams struct {
	Username    string       `json:"username"`
	Password    string       `json:"password"`
	Email       string       `json:"email"`
	SuperUser   sql.NullBool `json:"super_user"`
	IDMunicipio int64        `json:"id_municipio"`
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, insertUser,
		arg.Username,
		arg.Password,
		arg.Email,
		arg.SuperUser,
		arg.IDMunicipio,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.IDMunicipio,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.SuperUser,
		&i.CreatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, id_municipio, username, password, email, super_user, created_at FROM users
ORDER BY id
LIMIT $1
    OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.IDMunicipio,
			&i.Username,
			&i.Password,
			&i.Email,
			&i.SuperUser,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateToSuperUser = `-- name: UpdateToSuperUser :one
UPDATE users
SET
    "super_user" = $1
WHERE
        id = $2
RETURNING id, id_municipio, username, password, email, super_user, created_at
`

type UpdateToSuperUserParams struct {
	SuperUser sql.NullBool `json:"super_user"`
	ID        int64        `json:"id"`
}

func (q *Queries) UpdateToSuperUser(ctx context.Context, arg UpdateToSuperUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateToSuperUser, arg.SuperUser, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.IDMunicipio,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.SuperUser,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
    password = COALESCE($1, password)
WHERE
        id = $2
RETURNING id, id_municipio, username, password, email, super_user, created_at
`

type UpdateUserParams struct {
	Password sql.NullString `json:"password"`
	ID       int64          `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.Password, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.IDMunicipio,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.SuperUser,
		&i.CreatedAt,
	)
	return i, err
}
