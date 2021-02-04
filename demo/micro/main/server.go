package main

import (
	"context"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"microservice_monitor/demo/micro"
)

type HelloWorld struct{}

// 实现对应的 pb.go文件的方法才可以使用
// Call is a single request handler called via client.Call or the generated client code
func (h *HelloWorld) Call(ctx context.Context, req *micro.Request, rsp *micro.Response) error {
	logger.Info("Received HelloWorld.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	// Create service
	srv := service.New(
		service.Name("helloWorld"),
	)

	// Register Handler
	srv.Handle(new(HelloWorld))

	// Run the service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
