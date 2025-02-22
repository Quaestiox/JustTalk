package gapi

import (
	"database/sql"
	"github.com/Quaestiox/JustTalk_backend/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	userId := req.GetId()

	if authPayload.UserID != userId {
		return nil, status.Errorf(codes.PermissionDenied, "cannot delete other user's info")
	}

	err = server.store.DeleteUser(ctx, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to delete user: %s", err)
	}

	rep := &pb.DeleteUserResponse{
		Response: "deleted user successfully",
	}

	return rep, nil
}
