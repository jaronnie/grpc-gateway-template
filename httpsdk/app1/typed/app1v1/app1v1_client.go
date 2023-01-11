// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package app1v1

import (
	"github.com/jaronnie/grpc-gateway-template/httpsdk/app1/rest"
)

type App1v1Interface interface {
	RESTClient() rest.Interface

	HelloGetter
}

type App1v1Client struct {
	restClient rest.Interface
}

func (x *App1v1Client) RESTClient() rest.Interface {
	if x == nil {
		return nil
	}
	return x.restClient
}

func (x *App1v1Client) Hello() HelloInterface {
	return newHelloClient(x)
}

// NewForConfig creates a new App1v1Client for the given config.
func NewForConfig(x *rest.RESTClient) (*App1v1Client, error) {
	config := *x
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &App1v1Client{client}, nil
}
