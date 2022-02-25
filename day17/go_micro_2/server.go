package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/errors"
	pb "go_dev/day17/go_micro_2/proto"
	micro "github.com/micro/go-micro"
	"log"
)

type Example struct {}
type Foo struct {}
func (e *Example)Call(ctx context.Context, req *pb.CallRquest, res *pb.CallResponse) error{
	log.Print("收到Example.call请求")
	if len(req.Name)==0{
		return errors.BadRequest("go.micro.api.example","no name")
	}
	res.Message="Exmaple.call接收到了你的请求"+req.Name
	return nil
}
func (f *Foo)Bar(ctx context.Context, req *pb.EmptyRquest, res *pb.EmptyResponse) error {
	log.Print("收到Foo.Bar请求 :%#v",req)
	return nil
}
func main() {
	//1.new service
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
		micro.Version("latest"),
	)
	//2.initalise service
	service.Init()
	//3.Register struct as Example
	err := pb.RegisterExampleHandler(service.Server(), new(Example))
	if err!=nil{
		fmt.Println(err)
	}
	//4.register struct as Foo
	err =pb.RegisterFooHandler(service.Server(),new(Foo))
	if err!=nil{
		fmt.Println(err)
	}
	//5.run service
	if err :=service.Run();err!=nil{
		log.Fatal(err)
	}

}
