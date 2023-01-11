// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package fake

import (
	"context"

	"github.com/jaronnie/grpc-gateway-template/httpsdk/app1/pb/mypb/app1"
)

var (
	FakeReturnSayHello = &app1.Hello{}
)

type App1Getter interface {
	App1() App1Interface
}

type App1Interface interface {
	SayHello(ctx context.Context, param *app1.Hello) (*app1.Hello, error)
}

type FakeApp1 struct {
	Fake *FakeApp1v1
}

func (f *FakeApp1) SayHello(ctx context.Context, param *app1.Hello) (*app1.Hello, error) {
	return FakeReturnSayHello, nil
}
