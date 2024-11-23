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

-- name: CreateGame :one
INSERT INTO games (
    black, white
) values (
    ?, ?
) 
RETURNING *;
