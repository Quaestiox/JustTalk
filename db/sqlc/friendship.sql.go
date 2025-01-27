// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: friendship.sql

package db

import (
	"context"
)

const createFriendShip = `-- name: CreateFriendShip :one
INSERT INTO "friendship"(
                       "fromId",
                       "toId",
                       status
) VALUES (
    $1, $2, $3
) RETURNING id, "fromId", "toId", status, "createAt", "updateAt"
`

type CreateFriendShipParams struct {
	FromId int64 `json:"fromId"`
	ToId   int64 `json:"toId"`
	Status int16 `json:"status"`
}

func (q *Queries) CreateFriendShip(ctx context.Context, arg CreateFriendShipParams) (Friendship, error) {
	row := q.queryRow(ctx, q.createFriendShipStmt, createFriendShip, arg.FromId, arg.ToId, arg.Status)
	var i Friendship
	err := row.Scan(
		&i.ID,
		&i.FromId,
		&i.ToId,
		&i.Status,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const deleteFriendShip = `-- name: DeleteFriendShip :exec
DELETE FROM "friendship"
WHERE id = $1
`

func (q *Queries) DeleteFriendShip(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteFriendShipStmt, deleteFriendShip, id)
	return err
}

const getFriendShip = `-- name: GetFriendShip :one
SELECT id, "fromId", "toId", status, "createAt", "updateAt" FROM "friendship"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetFriendShip(ctx context.Context, id int64) (Friendship, error) {
	row := q.queryRow(ctx, q.getFriendShipStmt, getFriendShip, id)
	var i Friendship
	err := row.Scan(
		&i.ID,
		&i.FromId,
		&i.ToId,
		&i.Status,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const listFriendShip = `-- name: ListFriendShip :many
SELECT id, "fromId", "toId", status, "createAt", "updateAt" FROM "friendship"
WHERE ("fromId" = $1) OR ("toId" = $1)
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListFriendShipParams struct {
	FromId int64 `json:"fromId"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListFriendShip(ctx context.Context, arg ListFriendShipParams) ([]Friendship, error) {
	rows, err := q.query(ctx, q.listFriendShipStmt, listFriendShip, arg.FromId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Friendship{}
	for rows.Next() {
		var i Friendship
		if err := rows.Scan(
			&i.ID,
			&i.FromId,
			&i.ToId,
			&i.Status,
			&i.CreateAt,
			&i.UpdateAt,
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

const updateFriendShip = `-- name: UpdateFriendShip :one
UPDATE "friendship"
SET "status" = $1
WHERE id = $2
RETURNING id, "fromId", "toId", status, "createAt", "updateAt"
`

type UpdateFriendShipParams struct {
	Status int16 `json:"status"`
	ID     int64 `json:"id"`
}

func (q *Queries) UpdateFriendShip(ctx context.Context, arg UpdateFriendShipParams) (Friendship, error) {
	row := q.queryRow(ctx, q.updateFriendShipStmt, updateFriendShip, arg.Status, arg.ID)
	var i Friendship
	err := row.Scan(
		&i.ID,
		&i.FromId,
		&i.ToId,
		&i.Status,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}
