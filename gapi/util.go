package gapi

import (
	db "github.com/Quaestiox/JustTalk_backend/db/sqlc"
	"github.com/Quaestiox/JustTalk_backend/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.UserResponse {
	return &pb.UserResponse{
		Id:          user.ID,
		Name:        user.Name,
		Nickname:    user.Nickname,
		AvatarUrl:   user.AvatarUrl,
		FriendCount: user.FriendCount,
		Friends:     user.Friends,
		CreateAt:    timestamppb.New(user.CreateAt),
		UpdateAt:    timestamppb.New(user.UpdateAt),
	}
}
