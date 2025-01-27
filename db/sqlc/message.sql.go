// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: message.sql

package db

import (
	"context"
)

const createMessage = `-- name: CreateMessage :one
INSERT INTO "message"(
    "senderId",
    "receiverId",
    content
) VALUES (
             $1, $2, $3
) RETURNING id, "senderId", "receiverId", content, "sendAt"
`

type CreateMessageParams struct {
	SenderId   int64  `json:"senderId"`
	ReceiverId int64  `json:"receiverId"`
	Content    string `json:"content"`
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) (Message, error) {
	row := q.queryRow(ctx, q.createMessageStmt, createMessage, arg.SenderId, arg.ReceiverId, arg.Content)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.SenderId,
		&i.ReceiverId,
		&i.Content,
		&i.SendAt,
	)
	return i, err
}

const deleteMessage = `-- name: DeleteMessage :exec
DELETE FROM "message"
WHERE id = $1
`

func (q *Queries) DeleteMessage(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteMessageStmt, deleteMessage, id)
	return err
}

const getMessage = `-- name: GetMessage :one
SELECT id, "senderId", "receiverId", content, "sendAt" FROM "message"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMessage(ctx context.Context, id int64) (Message, error) {
	row := q.queryRow(ctx, q.getMessageStmt, getMessage, id)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.SenderId,
		&i.ReceiverId,
		&i.Content,
		&i.SendAt,
	)
	return i, err
}

const listMessage = `-- name: ListMessage :many
SELECT id, "senderId", "receiverId", content, "sendAt" FROM "message"
WHERE ("senderId" = $1) OR ("receiverId" = $1)
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListMessageParams struct {
	SenderId int64 `json:"senderId"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *Queries) ListMessage(ctx context.Context, arg ListMessageParams) ([]Message, error) {
	rows, err := q.query(ctx, q.listMessageStmt, listMessage, arg.SenderId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Message{}
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.SenderId,
			&i.ReceiverId,
			&i.Content,
			&i.SendAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
