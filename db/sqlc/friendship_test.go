package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateRandomFriendShip(t *testing.T) Friendship {
	arg := CreateFriendShipParams{
		FromId: 1,
		ToId:   2,
		Status: 0,
	}

	friendship, err := testQueries.CreateFriendShip(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, friendship)

	require.Equal(t, arg.FromId, friendship.FromId)
	require.Equal(t, arg.ToId, friendship.ToId)
	require.Equal(t, arg.Status, friendship.Status)

	require.NotZero(t, friendship.ID)
	require.NotZero(t, friendship.CreateAt)
	require.NotZero(t, friendship.UpdateAt)

	return friendship
}

func TestCreateFriendShip(t *testing.T) {
	CreateRandomFriendShip(t)
}
