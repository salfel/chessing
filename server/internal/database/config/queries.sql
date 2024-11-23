-- name: GetUser :one
SELECT * FROM users 
WHERE id = ? LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
    name
) values (
    ?
)
RETURNING *;
