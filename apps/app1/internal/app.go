package internal

import (
	"net/http"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/jaronnie/grpc-gateway-template/base/pkg/logx"
)

// Application base app framework interface
type Application interface {
	Init() error
	Run() error
	Exit() error
}

// App is main service
type App struct {
	GrpcServer *grpc.Server
	HTTPServer *http.Server
	GrpcPort   string
	HttpPort   string
}

// Init application init
func (app *App) Init() (err error) {
	logx.Logger()

	logx.Infof("app init")
	app.GrpcPort = viper.GetString("grpc.port")
	app.HttpPort = viper.GetString("http.port")
	return
}

// Run application run
func (app *App) Run() (err error) {
	logx.Infof("app run")
	go func() {
		_, err := app.grpcServer()
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		_, err := app.gatewayServer()
		if err != nil {
			panic(err)
		}
	}()

	return nil
}

// Exit application exit
func (app *App) Exit() (err error) {
	return nil
}
