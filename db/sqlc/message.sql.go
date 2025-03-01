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
    "sender_id",
    "receiver_id",
    content
) VALUES (
             $1, $2, $3
) RETURNING id, sender_id, receiver_id, content, send_at
`

type CreateMessageParams struct {
	SenderID   int64  `json:"sender_id"`
	ReceiverID int64  `json:"receiver_id"`
	Content    string `json:"content"`
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) (Message, error) {
	row := q.queryRow(ctx, q.createMessageStmt, createMessage, arg.SenderID, arg.ReceiverID, arg.Content)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.SenderID,
		&i.ReceiverID,
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
SELECT id, sender_id, receiver_id, content, send_at FROM "message"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMessage(ctx context.Context, id int64) (Message, error) {
	row := q.queryRow(ctx, q.getMessageStmt, getMessage, id)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.SenderID,
		&i.ReceiverID,
		&i.Content,
		&i.SendAt,
	)
	return i, err
}

const listMessage = `-- name: ListMessage :many
SELECT id, sender_id, receiver_id, content, send_at FROM "message"
WHERE ("sender_id" = $1) OR ("receiver_id" = $1)
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListMessageParams struct {
	SenderID int64 `json:"sender_id"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *Queries) ListMessage(ctx context.Context, arg ListMessageParams) ([]Message, error) {
	rows, err := q.query(ctx, q.listMessageStmt, listMessage, arg.SenderID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Message{}
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.SenderID,
			&i.ReceiverID,
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
