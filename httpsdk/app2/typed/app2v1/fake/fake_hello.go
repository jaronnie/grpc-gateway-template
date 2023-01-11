// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package fake

import (
	"context"

	"github.com/jaronnie/grpc-gateway-template/httpsdk/app2/pb/mypb/app2"
)

var (
	FakeReturnSayHello = &app2.Hello{}
)

type HelloGetter interface {
	Hello() HelloInterface
}

type HelloInterface interface {
	SayHello(ctx context.Context, param *app2.Hello) (*app2.Hello, error)
}

type FakeHello struct {
	Fake *FakeApp2v1
}

func (f *FakeHello) SayHello(ctx context.Context, param *app2.Hello) (*app2.Hello, error) {
	return FakeReturnSayHello, nil
}
