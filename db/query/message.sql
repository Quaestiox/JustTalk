-- name: CreateMessage :one
INSERT INTO "message"(
    "sender_id",
    "receiver_id",
    content
) VALUES (
             $1, $2, $3
) RETURNING *;

-- name: GetMessage :one
SELECT * FROM "message"
WHERE id = $1 LIMIT 1;

-- name: ListMessage :many
SELECT * FROM "message"
WHERE ("sender_id" = $1) OR ("receiver_id" = $1)
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: DeleteMessage :exec
DELETE FROM "message"
WHERE id = $1;