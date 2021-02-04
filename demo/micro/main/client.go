package main

import (
	"context"

	"github.com/micro/micro/v3/service/client"
	_ "github.com/micro/micro/v3/service/client"
	"microservice_monitor/demo/micro"
)

func clientS() {

	// 创建一个新的helloWorld服务客户端
	helloWorld := micro.NewHelloWorldService("helloWorld", client.DefaultClient)

	// 调用端点HelloWorld.Call
	rsp, err := helloWorld.Call(context.Background(), &micro.Request{Name: "Alice"})
	// TODO 全看业务
}
