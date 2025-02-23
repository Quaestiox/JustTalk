package gapi

import (
	"context"
	"github.com/Quaestiox/JustTalk_backend/pb"
	"github.com/Quaestiox/JustTalk_backend/rabbitmq"
	"github.com/Quaestiox/JustTalk_backend/util"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

func (server *Server) ReceiveMessage(ctx context.Context, req *pb.ReceiveMsgRequest) (*pb.ReceiveMsgResponse, error) {

	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	recId := authPayload.UserID
	senderId := req.SenderId
	rabbitmq := rabbitmq.NewRabbitMQSimple("chat" + strconv.Itoa(int(senderId)) + "-" + strconv.Itoa(int(recId)))
	msgsChan := rabbitmq.GetMsg()

	var msgs []*pb.Msg

	for msg := range msgsChan {
		msgId, err := strconv.Atoi(string(msg))
		util.Unwrap(err, "string to int failed.")
		getmsg, err := server.store.GetMessage(ctx, int64(msgId))
		message := &pb.Msg{
			Id:         getmsg.ID,
			SenderId:   getmsg.SenderID,
			ReceiverId: getmsg.ReceiverID,
			Content:    getmsg.Content,
			SendAt:     timestamppb.New(getmsg.SendAt),
		}
		msgs = append(msgs, message)
	}

	rsp := &pb.ReceiveMsgResponse{
		Message: msgs,
	}

	return rsp, nil
}
