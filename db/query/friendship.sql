-- name: CreateFriendShip :one
INSERT INTO "friendship"(
                       "fromId",
                       "toId",
                       status
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetFriendShip :one
SELECT * FROM "friendship"
WHERE id = $1 LIMIT 1;

-- name: ListFriendShip :many
SELECT * FROM "friendship"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateFriendShip :one
UPDATE "friendship"
SET "status" = $1
WHERE id = $2
RETURNING *;

-- name: DeleteFriendShip :exec
DELETE FROM "friendship"
WHERE id = $1;
