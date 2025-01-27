-- name: CreateMessage :one
INSERT INTO "message"(
    "senderId",
    "receiverId",
    content
) VALUES (
             $1, $2, $3
) RETURNING *;

-- name: GetMessage :one
SELECT * FROM "message"
WHERE id = $1 LIMIT 1;

-- name: ListMessage :many
SELECT * FROM "message"
WHERE ("senderId" = $1) OR ("receiverId" = $1)
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: DeleteMessage :exec
DELETE FROM "message"
WHERE id = $1;