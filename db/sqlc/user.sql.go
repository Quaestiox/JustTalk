// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"

	"github.com/lib/pq"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "user"(
                 name,
                 password,
                 nickname,
                 "avatarURL",
                 "friendCount",
                 friends
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id, name, password, nickname, "avatarURL", "friendCount", friends, "createAt", "updateAt"
`

type CreateUserParams struct {
	Name        string  `json:"name"`
	Password    string  `json:"password"`
	Nickname    string  `json:"nickname"`
	AvatarURL   string  `json:"avatarURL"`
	FriendCount int32   `json:"friendCount"`
	Friends     []int64 `json:"friends"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser,
		arg.Name,
		arg.Password,
		arg.Nickname,
		arg.AvatarURL,
		arg.FriendCount,
		pq.Array(arg.Friends),
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Password,
		&i.Nickname,
		&i.AvatarURL,
		&i.FriendCount,
		pq.Array(&i.Friends),
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, password, nickname, "avatarURL", "friendCount", friends, "createAt", "updateAt" FROM "user"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Password,
		&i.Nickname,
		&i.AvatarURL,
		&i.FriendCount,
		pq.Array(&i.Friends),
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const getUserForUpdate = `-- name: GetUserForUpdate :one
SELECT id, name, password, nickname, "avatarURL", "friendCount", friends, "createAt", "updateAt" FROM "user"
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetUserForUpdate(ctx context.Context, id int64) (User, error) {
	row := q.queryRow(ctx, q.getUserForUpdateStmt, getUserForUpdate, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Password,
		&i.Nickname,
		&i.AvatarURL,
		&i.FriendCount,
		pq.Array(&i.Friends),
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const listUser = `-- name: ListUser :many
SELECT id, name, password, nickname, "avatarURL", "friendCount", friends, "createAt", "updateAt" FROM "user"
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUserParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUser(ctx context.Context, arg ListUserParams) ([]User, error) {
	rows, err := q.query(ctx, q.listUserStmt, listUser, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Password,
			&i.Nickname,
			&i.AvatarURL,
			&i.FriendCount,
			pq.Array(&i.Friends),
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

const updateUser = `-- name: UpdateUser :one
UPDATE "user"
SET name = $1,password = $2, nickname = $3, "avatarURL" = $4, "friendCount" = $5, friends = $6
WHERE id = $7
RETURNING id, name, password, nickname, "avatarURL", "friendCount", friends, "createAt", "updateAt"
`

type UpdateUserParams struct {
	Name        string  `json:"name"`
	Password    string  `json:"password"`
	Nickname    string  `json:"nickname"`
	AvatarURL   string  `json:"avatarURL"`
	FriendCount int32   `json:"friendCount"`
	Friends     []int64 `json:"friends"`
	ID          int64   `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.queryRow(ctx, q.updateUserStmt, updateUser,
		arg.Name,
		arg.Password,
		arg.Nickname,
		arg.AvatarURL,
		arg.FriendCount,
		pq.Array(arg.Friends),
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Password,
		&i.Nickname,
		&i.AvatarURL,
		&i.FriendCount,
		pq.Array(&i.Friends),
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}
