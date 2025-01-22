package db

import (
	"context"
	"github.com/Quaestiox/JustTalk_backend/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Name:      util.RandomAccount(),
		Password:  util.RandomPassword(),
		Nickname:  util.RandomAccount(),
		AvatarURL: util.RandomPassword(),
		Friends:   []int64{},
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Nickname, user.Nickname)
	require.Equal(t, arg.AvatarURL, user.AvatarURL)
	require.Equal(t, arg.FriendCount, user.FriendCount)
	require.Equal(t, arg.Friends, user.Friends)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreateAt)
	require.NotZero(t, user.UpdateAt)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}
