-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)

VALUES (
	gen_random_uuid(),
	NOW(),
	NOW(),
	$1,
	$2
)

RETURNING *;


-- name: Reset :exec
DELETE FROM users;


-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;


-- name: SaveRefreshToken :one
INSERT INTO refresh_tokens (user_id, created_at, updated_at, token, expires_at)

VALUES (
	$1,
	NOW(),
	NOW(),
	$2,
	$3
)

RETURNING *;

-- name: GetRefreshToken :one
SELECT * FROM refresh_tokens WHERE token = $1 AND revoked_at IS NULL;


-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens SET revoked_at = NOW() WHERE token = $1;

-- name: UpdateUser :one
UPDATE users SET updated_at = NOW(), hashed_password = $2, email = $3 WHERE id = $1 RETURNING *;
