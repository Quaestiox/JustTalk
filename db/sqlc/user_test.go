package db

import (
	"context"
	"database/sql"
	"github.com/Quaestiox/JustTalk_backend/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateRandomUser(t *testing.T) User {
	hashedpd, err := util.HashPassword(util.RandomPassword())
	require.NoError(t, err)

	arg := CreateUserParams{
		Name:      util.RandomAccount(),
		Password:  hashedpd,
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

func TestGetUser(t *testing.T) {
	user := CreateRandomUser(t)
	userGet, err := testQueries.GetUser(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, userGet)

	require.Equal(t, user.ID, userGet.ID)
	require.Equal(t, user.Name, userGet.Name)
	require.Equal(t, user.Password, userGet.Password)
	require.Equal(t, user.Nickname, userGet.Nickname)
	require.Equal(t, user.AvatarURL, userGet.AvatarURL)
	require.Equal(t, user.FriendCount, userGet.FriendCount)
	require.Equal(t, user.Friends, userGet.Friends)
}

func TestUpdateUser(t *testing.T) {
	user := CreateRandomUser(t)

	friends := []int64{1, 2}

	arg := UpdateUserParams{
		ID:          user.ID,
		Name:        user.Name,
		Nickname:    util.RandomAccount(),
		AvatarURL:   util.RandomPassword(),
		Password:    util.RandomPassword(),
		Friends:     friends,
		FriendCount: int32(len(friends)),
	}

	userUp, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userUp)

	require.Equal(t, user.ID, userUp.ID)
	require.Equal(t, user.Name, userUp.Name)
	require.Equal(t, arg.Password, userUp.Password)
	require.Equal(t, arg.Nickname, userUp.Nickname)
	require.Equal(t, arg.AvatarURL, userUp.AvatarURL)
	require.Equal(t, arg.FriendCount, userUp.FriendCount)
	require.Equal(t, arg.Friends, userUp.Friends)

}

func TestDeleteUser(t *testing.T) {
	user := CreateRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)

	userDel, err := testQueries.GetUser(context.Background(), user.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, userDel)

}
