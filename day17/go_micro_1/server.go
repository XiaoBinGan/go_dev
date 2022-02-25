package main

import (
	"context"
	"fmt"
	micro "github.com/micro/go-micro/v2"
	pb "go_dev/day17/go_micro_1/proto"
	"log"
)

type Hello struct {}

func (h *Hello)Info(ctx context.Context, req *pb.InfoRequest, res *pb.InfoRespones) error  {
	res.Msg="Hello"+req.Username
	return  nil
}

func main() {
	//1.get server example
	service := micro.NewService(
		//set micro service name
		micro.Name("hello"),
	)
	//2.init service
	service.Init()
	//3.registry server
	err := pb.RegisterHelloHandler(service.Server(), new(Hello))
	if err!=nil{
		fmt.Println(err)
	}
	//4.start server
	if err := service.Run();err!=nil{
		log.Fatal(err)
	}

}
