package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFriendshipTx(t *testing.T) {
	store := NewStore(testDB)

	user1 := CreateRandomUser(t)
	user2 := CreateRandomUser(t)

	n := 5

	errChan := make(chan error)
	resultChan := make(chan FriendshipTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.FriendshipTx(context.Background(), FriendshipTxParams{
				FromUserId: user1.ID,
				ToUserId:   user2.ID,
			})

			errChan <- err
			resultChan <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errChan
		require.NoError(t, err)

		result := <-resultChan
		require.NotEmpty(t, result)

		friendship := result.Friendship
		require.NotEmpty(t, friendship)
		require.Equal(t, user1.ID, friendship.FromId)
		require.Equal(t, user2.ID, friendship.ToId)
		require.Equal(t, int16(0), friendship.Status)
		require.NotZero(t, friendship.ID)
		require.NotZero(t, friendship.CreateAt)
		require.NotZero(t, friendship.UpdateAt)

		_, err = store.GetFriendShip(context.Background(), friendship.ID)
		require.NoError(t, err)
	}

}
