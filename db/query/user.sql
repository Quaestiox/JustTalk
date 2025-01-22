-- name: CreateUser :one
INSERT INTO "user"(
                 name,
                 password,
                 nickname,
                 "avatarURL",
                 "friendCount",
                 friends
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUser :one
SELECT * FROM "user"
WHERE id = $1 LIMIT 1;

-- name: ListUser :many
SELECT * FROM "user"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUserAvatar :exec
UPDATE "user"
SET "avatarURL" = $1
WHERE id = $2;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;





