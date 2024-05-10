-- name: CreateUser :exec
INSERT INTO users (username, email, password) VALUES ($1, $2, $3);

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: CreatePastes :exec
INSERT INTO pastes (user_id, title, content, syntax, expiration_time, editable, slug) VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: GetPasteBySlug :one
SELECT * FROM pastes WHERE slug = $1 LIMIT 1;

-- name: GetPastesByUser :many
SELECT * FROM pastes WHERE user_id = $1;

-- name: DeletePaste :exec
DELETE FROM pastes WHERE id = $1;
