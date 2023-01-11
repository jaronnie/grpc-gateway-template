package api

import (
	"context"

	"github.com/jaronnie/grpc-gateway-template/base/mypb/app2"
)

type AppV2 struct{}

func (a AppV2) SayHello(ctx context.Context, hello *app2.Hello) (*app2.Hello, error) {
	return hello, nil
}
