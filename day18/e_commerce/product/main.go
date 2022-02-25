package main

import (
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_dev/day18/e_commerce/product/common"
	"go_dev/day18/e_commerce/product/domain/repository"
	service2 "go_dev/day18/e_commerce/product/domain/service"
	"go_dev/day18/e_commerce/product/handler"
	product "go_dev/day18/e_commerce/product/proto/product"
)

func main() {
	//config center
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err!=nil{
		logger.Error(err)
	}
	//registry center
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	//链路追踪
	tracer, io, err := common.NewTracer("go.micro.service.product", "localhost:6831")
	if err!=nil{
		logger.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(tracer)

	//数据库设置
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		log.Error(err)
	}
	defer db.Close()
	//禁止副表
	//在Gorm中，表名是结构体名的复数形式，列名是字段名的蛇形小写。即，如果有一个user表，那么如果你定义的结构体名为：User，gorm会默认表名为users而不是user。
    db.SingularTable(true)
	//初始化
    repository.NewProductRepository(db).InitTable()
	productDataService := service2.NewProductDataService(repository.NewProductRepository(db))
	// 设置服务
	service := micro.NewService(
		micro.Name("go.micro.service.product"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8085"),
		//添加注册中心&
		micro.Registry(consul),
		//绑定链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	// Initialise service
    service.Init()
	// Register Handler
    product.RegisterProductHandler(service.Server(),&handler.Product{ProductDateService: productDataService})
	// Run service
     if err :=service.Run();err!=nil{
     	log.Fatal(err)
	 }
}
