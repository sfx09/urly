-- name: CreateLink :one
INSERT INTO links (id, created_at, updated_at, full_link, short_link)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
