package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, txFunc func(queries *Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = txFunc(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type FriendshipTxParams struct {
	FromUserId int64 `json: "from_user_id"`
	ToUserId   int64 `json: "to_account_id"`
	Status     int16 `json: "status"`
}

type FriendshipTxResult struct {
	Friendship Friendship `json:"friendship"`
	FromUser   User       `json:"from_user"`
	ToUser     User       `json:"to_user"`
}

func (store *Store) FriendshipTx(ctx context.Context, arg FriendshipTxParams) (FriendshipTxResult, error) {
	var result FriendshipTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Friendship, err = q.CreateFriendShip(ctx, CreateFriendShipParams{
			FromId: arg.FromUserId,
			ToId:   arg.ToUserId,
			Status: arg.Status,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
