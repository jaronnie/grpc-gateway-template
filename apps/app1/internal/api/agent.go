package api

import (
	"context"

	"github.com/jaronnie/grpc-gateway-template/base/mypb"
)

type AppV1 struct{}

func (a AppV1) SayHello(ctx context.Context, hello *mypb.Hello) (*mypb.Hello, error) {
	return hello, nil
}
