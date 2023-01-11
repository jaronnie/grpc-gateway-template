// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package myservice

import (
	"github.com/jaronnie/grpc-gateway-template/httpsdk/myservice/rest"
	"github.com/jaronnie/grpc-gateway-template/httpsdk/myservice/typed"

	myservicev1 "github.com/jaronnie/grpc-gateway-template/httpsdk/myservice/typed/myservicev1"
)

type Interface interface {
	Direct() typed.DirectInterface

	Myservicev1() myservicev1.Myservicev1Interface
}

type Clientset struct {
	// direct client to request
	direct *typed.DirectClient

	myservicev1 *myservicev1.Myservicev1Client
}

func (x *Clientset) Direct() typed.DirectInterface {
	return x.direct
}

func (x *Clientset) Myservicev1() myservicev1.Myservicev1Interface {
	return x.myservicev1
}

func NewClientWithOptions(ops ...rest.Opt) (*Clientset, error) {
	c := &rest.RESTClient{}
	for _, op := range ops {
		if err := op(c); err != nil {
			return nil, err
		}
	}
	configShallowCopy := *c
	var cs Clientset
	var err error
	cs.direct, err = typed.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.myservicev1, err = myservicev1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return &cs, nil
}