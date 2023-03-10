package internal

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/jaronnie/grpc-gateway-template/apps/app2/internal/api"
	"github.com/jaronnie/grpc-gateway-template/base/mypb"
	"github.com/jaronnie/grpc-gateway-template/base/pkg/logx"
)

func (app *App) grpcServer() (s *grpc.Server, err error) {
	logx.Infof("start grpc server 0.0.0.0:%s\n", app.GrpcPort)
	listen, err := net.Listen("tcp", "0.0.0.0:"+app.GrpcPort)
	if err != nil {
		panic(err)
	}

	s = grpc.NewServer()
	// Register reflection service on gRPC server.
	reflection.Register(s)

	appV2 := &api.AppV2{}
	mypb.RegisterApp2Server(s, appV2)

	app.GrpcServer = s

	return s, s.Serve(listen)
}

func (app *App) gatewayServer() (s *http.Server, err error) {
	logx.Infof("start gateway server 0.0.0.0:%s\n", app.HttpPort)
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := mypb.RegisterApp2HandlerFromEndpoint(context.Background(), mux, "0.0.0.0:"+app.GrpcPort, opts); err != nil {
		return nil, err
	}

	s = &http.Server{
		Addr:    "0.0.0.0:" + app.HttpPort,
		Handler: mux,
	}

	app.HTTPServer = s

	return s, s.ListenAndServe()
}
