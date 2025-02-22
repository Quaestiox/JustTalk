package gapi

import (
	"database/sql"
	"github.com/Quaestiox/JustTalk_backend/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	userId := req.GetId()
	user, err := server.store.GetUser(ctx, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to delete user: %s", err)
	}

	userArg := &pb.UserResponse{
		Id:          user.ID,
		Name:        user.Name,
		Nickname:    user.Nickname,
		AvatarUrl:   user.AvatarUrl,
		FriendCount: user.FriendCount,
		Friends:     user.Friends,
		CreateAt:    timestamppb.New(user.CreateAt),
		UpdateAt:    timestamppb.New(user.UpdateAt),
	}

	rep := &pb.GetUserInfoResponse{
		User: userArg,
	}

	return rep, nil
}
