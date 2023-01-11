package api

import (
	"context"

	"github.com/jaronnie/grpc-gateway-template/base/mypb/app1"
)

type AppV1 struct{}

func (a AppV1) SayHello(ctx context.Context, hello *app1.Hello) (*app1.Hello, error) {
	return hello, nil
}
