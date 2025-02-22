package gapi

import (
	"context"
	"fmt"
	db "github.com/Quaestiox/JustTalk_backend/db/sqlc"
	"github.com/Quaestiox/JustTalk_backend/pb"
	"github.com/Quaestiox/JustTalk_backend/util"
)

type Server struct {
	pb.UnimplementedJustTalkServer
	config     util.Config
	store      *db.Store
	tokenMaker util.Maker
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := util.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

	return server, nil
}

func (server *Server) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{Answer: "hello " + req.GetName()}, nil
}
