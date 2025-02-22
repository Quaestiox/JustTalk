package gapi

import (
	"context"
	"database/sql"
	"github.com/Quaestiox/JustTalk_backend/pb"
	"github.com/Quaestiox/JustTalk_backend/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := server.store.GetUserByName(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to find user")
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "incorrect password")
	}

	accessToken, _, err := server.tokenMaker.CreateToken(
		user.ID,
		user.Name,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token")
	}

	rsp := &pb.LoginUserResponse{
		AccessToken: accessToken,
		User:        convertUser(user),
	}

	return rsp, nil
}
