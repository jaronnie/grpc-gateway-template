// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package fake

import (
	"github.com/jaronnie/grpc-gateway-template/httpsdk/app2/typed"
	"github.com/jaronnie/grpc-gateway-template/httpsdk/app2/typed/fake"

	app2v1 "github.com/jaronnie/grpc-gateway-template/httpsdk/app2/typed/app2v1"
	fakeapp2v1 "github.com/jaronnie/grpc-gateway-template/httpsdk/app2/typed/app2v1/fake"
)

type Clientset struct{}

func (f *Clientset) Direct() typed.DirectInterface {
	return &fake.FakeDirect{}
}

func (f *Clientset) App2v1() app2v1.App2v1Interface {
	return &fakeapp2v1.FakeApp2v1{}
}
