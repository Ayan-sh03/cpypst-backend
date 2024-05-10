// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package generated

import (
	"context"
	"database/sql"
)

const createPastes = `-- name: CreatePastes :exec
INSERT INTO pastes (user_id, title, content, syntax, expiration_time, editable, slug) VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type CreatePastesParams struct {
	UserID         sql.NullInt32
	Title          sql.NullString
	Content        string
	Syntax         sql.NullString
	ExpirationTime sql.NullTime
	Editable       sql.NullBool
	Slug           string
}

func (q *Queries) CreatePastes(ctx context.Context, arg CreatePastesParams) error {
	_, err := q.db.ExecContext(ctx, createPastes,
		arg.UserID,
		arg.Title,
		arg.Content,
		arg.Syntax,
		arg.ExpirationTime,
		arg.Editable,
		arg.Slug,
	)
	return err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO users (username, email, password) VALUES ($1, $2, $3)
`

type CreateUserParams struct {
	Username string
	Email    string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.Username, arg.Email, arg.Password)
	return err
}

const deletePaste = `-- name: DeletePaste :exec
DELETE FROM pastes WHERE id = $1
`

func (q *Queries) DeletePaste(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deletePaste, id)
	return err
}

const getPasteBySlug = `-- name: GetPasteBySlug :one
SELECT id, user_id, title, content, syntax, expiration_time, editable, slug, created_at, updated_at FROM pastes WHERE slug = $1 LIMIT 1
`

func (q *Queries) GetPasteBySlug(ctx context.Context, slug string) (Paste, error) {
	row := q.db.QueryRowContext(ctx, getPasteBySlug, slug)
	var i Paste
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Content,
		&i.Syntax,
		&i.ExpirationTime,
		&i.Editable,
		&i.Slug,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPastesByUser = `-- name: GetPastesByUser :many
SELECT id, user_id, title, content, syntax, expiration_time, editable, slug, created_at, updated_at FROM pastes WHERE user_id = $1
`

func (q *Queries) GetPastesByUser(ctx context.Context, userID sql.NullInt32) ([]Paste, error) {
	rows, err := q.db.QueryContext(ctx, getPastesByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Paste
	for rows.Next() {
		var i Paste
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Content,
			&i.Syntax,
			&i.ExpirationTime,
			&i.Editable,
			&i.Slug,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, email, password, created_at FROM users WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}