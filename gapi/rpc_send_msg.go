package gapi

import (
	"context"
	db "github.com/Quaestiox/JustTalk_backend/db/sqlc"
	"github.com/Quaestiox/JustTalk_backend/pb"
	"github.com/Quaestiox/JustTalk_backend/rabbitmq"
	"github.com/Quaestiox/JustTalk_backend/util"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strconv"
)

func (server *Server) SendMessage(ctx context.Context, req *pb.CreateMsgRequest) (*pb.CreateMsgResponse, error) {

	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	content := req.GetContent()
	secretMsg, err := util.EncryptAES(content)
	util.Unwrap(err, "cannot encrypt message")
	arg := db.CreateMessageParams{
		SenderID:   authPayload.UserID,
		ReceiverID: req.GetReceiverId(),
		Content:    secretMsg,
	}

	msg, err := server.store.CreateMessage(ctx, arg)
	util.Unwrap(err, "create message failed.")

	rabbitmq := rabbitmq.NewRabbitMQSimple("chat" + strconv.Itoa(int(msg.SenderID)) + "-" + strconv.Itoa(int(msg.ReceiverID)))
	rabbitmq.PublishSimple(strconv.Itoa(int(msg.ID)))
	log.Printf("User[%d] sended successfully.\n", msg.SenderID)

	rspMsg := &pb.Msg{
		Id:         msg.ID,
		SenderId:   msg.SenderID,
		ReceiverId: msg.ReceiverID,
		Content:    msg.Content,
		SendAt:     timestamppb.New(msg.SendAt),
	}
	rsp := &pb.CreateMsgResponse{
		Message: rspMsg,
	}

	return rsp, nil
}
