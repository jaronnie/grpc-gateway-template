package main

import (
	"context"
	"fmt"
	"github.com/jaronnie/grpc-gateway-template/pkg/myservicesdk"
	"github.com/jaronnie/grpc-gateway-template/pkg/myservicesdk/pb/mypb"
	"github.com/jaronnie/grpc-gateway-template/pkg/myservicesdk/rest"
	"log"
)

func main() {
	var clientSet myservicesdk.Interface
	var err error
	clientSet, err = myservicesdk.NewClientWithOptions(
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
	hello, err := clientSet.Myservicev1().Hello().SayHello(context.Background(), &mypb.Hello{
		Message: "hello jaronnie!",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hello)
}
