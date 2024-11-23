-- name: GetUser :one
SELECT * FROM users 
WHERE username = ? LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    username, password
) values (
    ?, ?
)
RETURNING *;

-- name: CreateGame :exec
INSERT INTO games (
    black, white
) values (
    ?, ?
) 
RETURNING *;

-- name: JoinGame :exec
UPDATE games 
SET white = sqlc.arg('white'), black = sqlc.arg('black')
WHERE id = sqlc.arg('id');

-- name: GetGame :one
SELECT * FROM games
WHERE id = ?;
