package main

import (
	"grpc/chat"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to connect on port 9k: %v", err)
	}

	s := chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve on 9k: %v", err)
	}

}
