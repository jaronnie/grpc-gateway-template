package main

import (
	"context"
	"fmt"
	"github.com/jaronnie/grpc-gateway-template/base/mypb"
	"google.golang.org/grpc"
	"log"
)

var address = "localhost:9603"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := mypb.NewApp1Client(conn)

	app, err := client.SayHello(context.Background(), &mypb.Hello{
		Message: "Hello jaronnie!",
	})
	if err != nil {
		log.Fatalf("say hello meet error: %v", err)
	}
	fmt.Println(app)
}
