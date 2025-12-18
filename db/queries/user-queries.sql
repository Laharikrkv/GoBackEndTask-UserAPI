-- name: CreateUser :one
INSERT INTO users (
  name, dob
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users SET name = $2, dob = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: GetUser :many
SELECT * FROM users;