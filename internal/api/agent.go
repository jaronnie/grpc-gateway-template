package api

import (
	"context"

	"github.com/jaronnie/grpc-gateway-template/mypb"
)

type MyServiceV1 struct{}

func (a MyServiceV1) SayHello(ctx context.Context, hello *mypb.Hello) (*mypb.Hello, error) {
	return hello, nil
}
