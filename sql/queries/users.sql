-- name: CreateUser :one
INSERT INTO users (
        username,
        password_hash
    )
VALUES ($1, $2)
RETURNING *;
-- name: GetUserByUsername :one
SELECT users.*,
    coalesce(
        json_agg(victims) filter (
            where victims is not null
        ),
        '[]'::json
    )::json victims
from users
    Left JOIN victims ON users.id = victims.user_id
where users.username = $1
GROUP BY users.id;