package gapi

import (
	"context"
	"github.com/Quaestiox/JustTalk_backend/pb"
	"github.com/Quaestiox/JustTalk_backend/rabbitmq"
	"github.com/Quaestiox/JustTalk_backend/util"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
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
	msg, ok := rabbitmq.GetOneMsg()
	if !ok {
		message := &pb.Msg{
			Id:         0,
			SenderId:   0,
			ReceiverId: 0,
			Content:    "",
			SendAt:     timestamppb.Now(),
		}
		rsp := &pb.ReceiveMsgResponse{
			Message: message,
		}
		log.Printf("There is no message for User[%d].\n", recId)
		return rsp, nil
	}

	log.Printf("User[%d] receive one message.\n", recId)
	msgId, err := strconv.Atoi(string(msg))
	util.Unwrap(err, "string to int failed.")
	getmsg, err := server.store.GetMessage(ctx, int64(msgId))
	util.Unwrap(err, "cannot get message")

	encryptContent := getmsg.Content
	realContent, err := util.DecryptAES(encryptContent)
	util.Unwrap(err, "cannot decrypted content.")

	message := &pb.Msg{
		Id:         getmsg.ID,
		SenderId:   getmsg.SenderID,
		ReceiverId: getmsg.ReceiverID,
		Content:    realContent,
		SendAt:     timestamppb.New(getmsg.SendAt),
	}
	rsp := &pb.ReceiveMsgResponse{
		Message: message,
	}

	return rsp, nil
}
