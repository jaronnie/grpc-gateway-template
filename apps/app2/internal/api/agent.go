package api

import (
	"context"

	"github.com/jaronnie/grpc-gateway-template/base/mypb"
)

type AppV2 struct{}

func (a AppV2) SayHello(ctx context.Context, hello *mypb.Hello) (*mypb.Hello, error) {
	return hello, nil
}
