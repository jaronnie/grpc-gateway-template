// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package fake

import (
	"github.com/jaronnie/grpc-gateway-template/httpsdk/myservice/rest"
	"github.com/jaronnie/grpc-gateway-template/httpsdk/myservice/typed/myservicev1"
)

type FakeMyservicev1 struct{}

func (f *FakeMyservicev1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (f *FakeMyservicev1) App1() myservicev1.App1Interface {
	return &FakeApp1{Fake: f}
}

func (f *FakeMyservicev1) App2() myservicev1.App2Interface {
	return &FakeApp2{Fake: f}
}