// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package database

import (
	"context"
	"database/sql"
)

const createGame = `-- name: CreateGame :exec
INSERT INTO games (
    black, white
) values (
    ?, ?
) 
RETURNING id, white, black
`

type CreateGameParams struct {
	Black sql.NullInt64
	White sql.NullInt64
}

func (q *Queries) CreateGame(ctx context.Context, arg CreateGameParams) error {
	_, err := q.db.ExecContext(ctx, createGame, arg.Black, arg.White)
	return err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username, password
) values (
    ?, ?
)
RETURNING id, username, password
`

type CreateUserParams struct {
	Username string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.Password)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.Password)
	return i, err
}

const getGame = `-- name: GetGame :one
SELECT id, white, black FROM games
WHERE id = ?
`

func (q *Queries) GetGame(ctx context.Context, id int64) (Game, error) {
	row := q.db.QueryRowContext(ctx, getGame, id)
	var i Game
	err := row.Scan(&i.ID, &i.White, &i.Black)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, password FROM users 
WHERE username = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.Password)
	return i, err
}

const joinGame = `-- name: JoinGame :exec
UPDATE games 
SET white = ?1, black = ?2
WHERE id = ?3
`

type JoinGameParams struct {
	White sql.NullInt64
	Black sql.NullInt64
	ID    int64
}

func (q *Queries) JoinGame(ctx context.Context, arg JoinGameParams) error {
	_, err := q.db.ExecContext(ctx, joinGame, arg.White, arg.Black, arg.ID)
	return err
}
