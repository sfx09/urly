-- name: CreateLink :one
INSERT INTO links (id, created_at, updated_at, full_link, short_link)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetByShortLink :one
SELECT * FROM links
WHERE short_link = $1;

-- name: UpdateLinkCounter :exec
UPDATE links
SET counter = counter + 1, updated_at = $2
WHERE short_link = $1;

-- name: DeleteExpiredLinks :exec
DELETE FROM links
WHERE AGE(updated_at) > INTERVAL '1 hours';
