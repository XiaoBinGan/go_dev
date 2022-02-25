package main

import (
	//"github.com/micro/go-micro"go-micro
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul"

	//consul "github.com/micro/go-plugins/registry/consul"
	//"go_dev/day18/go-micro-grpc/ServiceImpl"
	//"go_dev/day18/go-micro-grpc/Services"

 	//consul "github.com/micro/go-plugins/registry/consul"
)

func main()  {
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
		//registry.Addrs(),
	)
	service := micro.NewService(
		micro.Name("prodservice"),
		micro.Address(":8011"),
		micro.Registry(consulReg),
	)
	service.Init()
	//Services.RegisterProdServiceHandler(service.Server(),new(ServiceImpl.ProdService))
	service.Run()
}


