package main

import (
	"context"
	"fmt"
	pb "go_dev/day17/gRPC_1/proto"
	"google.golang.org/grpc"
	"net"
)

/**
1.需要监听
2.需要实例化GRPC服务端
3.在GRPC商注册微服务
4.启动服务端
*/


//定义空接口
type UserInfoService struct {}
var u = UserInfoService{}
//实现方法
func(s *UserInfoService)GetUserInfo(ctx context.Context,req *pb.UserRequest)(resp *pb.UserRespones,err error){
	//通过用户名查询用户信息
	name :=req.Name
	//数据库查询用户信息
	if name  =="Alben"{
		resp =&pb.UserRespones{
				Id: 1,
				Name: "Alben",
				Age: 18,
				Hobby: []string{"sing","run"},

		}
	}
	return
}
func main()  {
	addr :="127.0.0.1:8080"
	listener, err := net.Listen("tcp", addr)
	if err!=nil{
		fmt.Printf("监听异常:%s\n",err)
		return
	}
	fmt.Printf("监听端口:%s\n",addr)
	//2.实例化GRPC
	s :=grpc.NewServer()
	//3.在grpc上注册微服务
	pb.RegisterUserInfoServiceServer(s,&u)
	//4.启动服务端
	s.Serve(listener)
}