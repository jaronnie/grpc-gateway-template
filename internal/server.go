package internal

import (
	"context"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	adapter "github.com/gwatts/gin-adapter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/jaronnie/grpc-gateway-template/internal/api"
	"github.com/jaronnie/grpc-gateway-template/internal/api/router"
	"github.com/jaronnie/grpc-gateway-template/mypb"
	"github.com/jaronnie/grpc-gateway-template/pkg/logx"
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

	myServiceV1 := &api.MyServiceV1{}
	mypb.RegisterMyServiceServer(s, myServiceV1)

	app.GrpcServer = s

	return s, s.Serve(listen)
}

// gatewayServer
// Deprecated instead of httpServer
func (app *App) gatewayServer() (s *http.Server, err error) {
	logx.Infof("start gateway server 0.0.0.0:%s\n", app.HttpPort)
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := mypb.RegisterMyServiceHandlerFromEndpoint(context.Background(), mux, "0.0.0.0:"+app.GrpcPort, opts); err != nil {
		return nil, err
	}

	s = &http.Server{
		Addr:    "0.0.0.0:" + app.HttpPort,
		Handler: mux,
	}

	app.HTTPServer = s

	return s, s.ListenAndServe()
}

func (app *App) httpServer() (s *http.Server, err error) {
	logx.Infof("start http server 0.0.0.0:%s\n", app.HttpPort)
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := mypb.RegisterMyServiceHandlerFromEndpoint(context.Background(), mux, "0.0.0.0:"+app.GrpcPort, opts); err != nil {
		return nil, err
	}

	g := gin.New()

	// wrap grpc gateway handler
	handler := adapter.Wrap(func(h http.Handler) http.Handler {
		return mux
	})

	// load gin handler
	g = router.Load(g)

	g.Use(handler)

	s = &http.Server{
		Addr:    "0.0.0.0:" + app.HttpPort,
		Handler: g,
	}

	app.HTTPServer = s

	return s, s.ListenAndServe()
}
