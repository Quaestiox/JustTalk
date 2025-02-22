package gapi

import (
	"context"
	db "github.com/Quaestiox/JustTalk_backend/db/sqlc"
	"github.com/Quaestiox/JustTalk_backend/pb"
	"github.com/Quaestiox/JustTalk_backend/util"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPd, err := util.HashPassword(req.GetPassword())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	arg := db.CreateUserParams{
		Name:      req.GetName(),
		Password:  hashedPd,
		Nickname:  req.GetNickname(),
		AvatarUrl: req.GetAvatarUrl(),
		Friends:   req.GetFriends(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	rsp := &pb.CreateUserResponse{
		Id:          user.ID,
		Name:        user.Name,
		Nickname:    user.Nickname,
		AvatarUrl:   user.AvatarUrl,
		FriendCount: user.FriendCount,
		Friends:     user.Friends,
		CreateAt:    timestamppb.New(user.CreateAt),
		UpdateAt:    timestamppb.New(user.UpdateAt),
	}

	return rsp, nil
}
