-- name: CreateUser :one
INSERT INTO "user"(
                 name,
                 password,
                 nickname,
                 "avatar_url",
                 "friend_count",
                 friends
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUser :one
SELECT * FROM "user"
WHERE id = $1 LIMIT 1;

-- name: GetUserByName :one
SELECT * FROM "user"
WHERE name = $1 LIMIT 1;


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
SET name = COALESCE(sqlc.narg(name), name),
    password = COALESCE(sqlc.narg(password), password),
    nickname = COALESCE(sqlc.narg(nickname), nickname),
    avatar_url = COALESCE(sqlc.narg(avatar_url), avatar_url),
    friend_count = COALESCE(sqlc.narg(friend_count), friend_count),
    friends = COALESCE(sqlc.narg(friends), friends)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;





