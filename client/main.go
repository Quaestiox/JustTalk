package main

import (
	"context"
	"fmt"
	"github.com/Quaestiox/JustTalk_backend/pb"
	"google.golang.org/api/option"
	"google.golang.org/api/transport/grpc"
	"log"
)

func main() {
	conn, err := grpc.DialInsecure(context.Background(), option.WithEndpoint("localhost:9101"), option.WithoutAuthentication())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewJustTalkClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.SayHelloRequest{
		Name: "guest1415",
	})

	if err != nil {
		log.Fatalf("err:%v", err)
	}

	fmt.Println("answer:", resp.GetAnswer())
}
