-- name: CreateAPIToken :one
INSERT INTO api_tokens (
    user_id,
    name,
    token_hash,
    token_prefix,
    permissions,
    expires_at
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetAPIToken :one
SELECT * FROM api_tokens
WHERE id = $1 LIMIT 1;

-- name: GetAPITokenByHash :one
SELECT * FROM api_tokens
WHERE token_hash = $1 LIMIT 1;

-- name: ListAPITokensByUser :many
SELECT * FROM api_tokens
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: DeleteAPIToken :exec
DELETE FROM api_tokens
WHERE id = $1;

-- name: UpdateAPITokenLastUsedAt :exec
UPDATE api_tokens
SET last_used_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: CountAPITokensByUser :one
SELECT COUNT(*) FROM api_tokens
WHERE user_id = $1;
