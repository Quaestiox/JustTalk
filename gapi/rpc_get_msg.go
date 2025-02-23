package gapi

import (
	"github.com/Quaestiox/JustTalk_backend/pb"
	"github.com/Quaestiox/JustTalk_backend/util"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetMessage(ctx context.Context, req *pb.GetMsgRequest) (*pb.GetMsgResponse, error) {

	msg, err := server.store.GetMessage(ctx, req.GetId())
	util.Unwrap(err, "cannot get message")

	msgRsp := &pb.Msg{
		Id:         msg.ID,
		SenderId:   msg.SenderID,
		ReceiverId: msg.ReceiverID,
		Content:    msg.Content,
		SendAt:     timestamppb.New(msg.SendAt),
	}

	rsp := &pb.GetMsgResponse{
		Message: msgRsp,
	}

	return rsp, nil
}
