package main

import (
	"context"
	"fmt"
	pb "go_dev/day17/gRPC_1/proto"
	"google.golang.org/grpc"
)
/**
1.Dial
2.construct grpc
3.request params
4.calling interface
 */
func main()  {
	//1.
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err!=nil{
		fmt.Printf("连接异常")
	}
	defer conn.Close()
	//2.
	client := pb.NewUserInfoServiceClient(conn)
	//3.
	req := new(pb.UserRequest)
	req.Name="Alben"
	//4.
	respones, err := client.GetUserInfo(context.Background(), req)
	if err!=nil{
		fmt.Println("响应异常 %s\n",err)
	}
	fmt.Printf("响应结果:%v\n",respones)
}