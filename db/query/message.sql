-- name: CreateMessage :one
INSERT INTO "message"(
    "senderId",
    "receiverId",
    context
) VALUES (
             $1, $2, $3
) RETURNING *;

-- name: GetMessage :one
SELECT * FROM "message"
WHERE id = $1 LIMIT 1;

-- name: ListMessage :many
SELECT * FROM "message"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteMessage :exec
DELETE FROM "message"
WHERE id = $1;