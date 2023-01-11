package main

import (
	"context"
	"fmt"
	"github.com/jaronnie/grpc-gateway-template/httpsdk/app1"
	app1pb "github.com/jaronnie/grpc-gateway-template/httpsdk/app1/pb/mypb/app1"
	"github.com/jaronnie/grpc-gateway-template/httpsdk/app1/rest"
	"log"
)

func main() {
	var clientSet app1.Interface
	var err error
	clientSet, err = app1.NewClientWithOptions(
		rest.WithProtocol("http"),
		rest.WithAddr("127.0.0.1"),
		rest.WithPort("8090"),
		rest.WithHeaders(map[string][]string{
			"Content-Type": {"application/json"},
		}),
	)

	if err != nil {
		log.Fatal(err)
	}
	hello, err := clientSet.App1v1().Hello().SayHello(context.Background(), &app1pb.Hello{
		Message: "hello jaronnie!",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hello)
}
