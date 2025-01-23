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

-- name: GetUserForUpdate :one
SELECT * FROM "user"
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE ;

-- name: ListUser :many
SELECT * FROM "user"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE "user"
SET name = $1,password = $2, nickname = $3, "avatarURL" = $4, "friendCount" = $5, friends = $6
WHERE id = $7
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;





