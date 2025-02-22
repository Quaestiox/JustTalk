package gapi

import (
	"context"
	"database/sql"
	db "github.com/Quaestiox/JustTalk_backend/db/sqlc"
	"github.com/Quaestiox/JustTalk_backend/pb"
	"github.com/Quaestiox/JustTalk_backend/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {

	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if authPayload.UserID != req.GetId() {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update other user's info")
	}

	arg := db.UpdateUserParams{
		ID: req.GetId(),
		Name: sql.NullString{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},
		Nickname: sql.NullString{
			String: req.GetNickname(),
			Valid:  req.Nickname != nil,
		},
		AvatarUrl: sql.NullString{
			String: req.GetAvatarUrl(),
			Valid:  req.AvatarUrl != nil,
		},
	}

	if req.Friends != nil {
		arg.Friends = req.GetFriends()
		count := int32(len(arg.Friends))
		arg.FriendCount = sql.NullInt32{
			Int32: count,
			Valid: true,
		}
	}

	if req.Password != nil {
		hashedPd, err := util.HashPassword(req.GetPassword())

		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
		}

		arg.Password = sql.NullString{
			String: hashedPd,
			Valid:  true,
		}
	}

	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update user: %s", err)
	}

	rsp := &pb.UpdateUserResponse{
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
