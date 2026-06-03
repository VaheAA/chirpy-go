-- name: CreateChirp :one
INSERT INTO chirps (id, user_id, created_at, updated_at, body)

VALUES (
	gen_random_uuid(),
	$1,
	NOW(),
	NOW(),
	$2
)

RETURNING *;


-- name: ResetChirps :exec
DELETE FROM chirps;

-- name: GetChirps :many
SELECT * FROM chirps ORDER BY created_at ASC;

-- name: GetChirp :one
SELECT * FROM chirps WHERE id = $1;

-- name: DeleteChirp :exec
DELETE FROM chirps WHERE id = $1 AND user_id = $2;
