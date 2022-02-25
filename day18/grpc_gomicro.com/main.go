package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	greeter "go_dev/day18/grpc_gomicro.com/proto"
	"log"
)

type Greeter struct {}

func (g *Greeter)Hello(ctx context.Context, req *greeter.Request, res *greeter.Response)error  {
    log.Fatalln("接收 Greeter.Request 请求")
    res.Msg="你好"+req.Name
	return nil
}
func main()  {
	var service = micro.NewService(
		micro.Name("greeter"),
	)
	service.Init()
	greeter.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err := service.Run(); err != nil {
		log.Fatalln(err)
	}
}