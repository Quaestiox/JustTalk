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
WHERE ("fromId" = $1) OR ("toId" = $1)
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateFriendShip :one
UPDATE "friendship"
SET "status" = $1
WHERE id = $2
RETURNING *;

-- name: DeleteFriendShip :exec
DELETE FROM "friendship"
WHERE id = $1;
