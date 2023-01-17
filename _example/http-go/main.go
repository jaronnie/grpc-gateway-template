package main

import (
	"context"
	"fmt"
	"github.com/jaronnie/grpc-gateway-template/httpsdk/myservice"
	"github.com/jaronnie/grpc-gateway-template/httpsdk/myservice/pb/mypb"
	"github.com/jaronnie/grpc-gateway-template/httpsdk/myservice/rest"
	"log"
)

func main() {
	var clientSet myservice.Interface
	var err error
	clientSet, err = myservice.NewClientWithOptions(
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
	hello, err := clientSet.Myservicev1().App1().SayHello(context.Background(), &mypb.Hello{
		Message: "hello jaronnie!",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hello)
}
