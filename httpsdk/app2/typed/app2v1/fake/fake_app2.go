// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package fake

import (
	"context"

	"github.com/jaronnie/grpc-gateway-template/httpsdk/app2/pb/mypb/app2"
)

var (
	FakeReturnSayHello = &app2.Hello{}
)

type App2Getter interface {
	App2() App2Interface
}

type App2Interface interface {
	SayHello(ctx context.Context, param *app2.Hello) (*app2.Hello, error)
}

type FakeApp2 struct {
	Fake *FakeApp2v1
}

func (f *FakeApp2) SayHello(ctx context.Context, param *app2.Hello) (*app2.Hello, error) {
	return FakeReturnSayHello, nil
}
