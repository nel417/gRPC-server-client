package main

import (
	"grpc/chat"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	message := chat.Message{
		Body: "hello from client",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", response.Body)
}
