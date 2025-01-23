package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateMessage(t *testing.T) {
	arg := CreateMessageParams{
		SenderId:   1,
		ReceiverId: 2,
		Context:    "Hello!",
	}

	message, err := testQueries.CreateMessage(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, message.SenderId, arg.SenderId)
	require.Equal(t, message.ReceiverId, arg.ReceiverId)
	require.Equal(t, message.Context, arg.Context)

	require.NotZero(t, message.ID)
	require.NotZero(t, message.SendAt)

}
