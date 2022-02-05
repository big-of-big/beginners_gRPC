package main

import (
	"log"

	"github.com/big-of-big/beginners_gRPC/chat"
	"github.com/k0kubun/pp"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn // connection
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hey!",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error occurred when calling SayHello: %s", err)
	}

	pp.Println(response)
	log.Printf("Response from Server: %s", response.Body)
}
