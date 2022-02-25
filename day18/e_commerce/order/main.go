package main

import (
	"fmt"
	"github.com/XiaoBinGan/common"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	"github.com/opentracing/opentracing-go"
	"go_dev/day18/e_commerce/order/domain/repository"
	service2 "go_dev/day18/e_commerce/order/domain/service"
	"go_dev/day18/e_commerce/order/handler"
	order "go_dev/day18/e_commerce/order/proto/order"
)
var (
	QPS =100
)
func main()  {
	//config center
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err!=nil{
		log.Error(err)
	}
	//registry center
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	//jaeger tracer
	newTracer, io, err := common.NewTracer("go.micro.service.order", "localhost:6831")
	if err!=nil{
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(newTracer)
	//mysql config buffer
	mysqlConfig := common.GetMysqlFromConsul(consulConfig,"mysql")
	fmt.Printf("%+#v",mysqlConfig)
	//DB connect
	db, err := gorm.Open("mysql", mysqlConfig.User+":"+mysqlConfig.Pwd+"@/"+mysqlConfig.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		log.Error(err)
		fmt.Println("123123123231123")
	}
	defer db.Close()
	db.SingularTable(true)
	//DB model init
	//repository.NewOrderRepository(db).InitTable() //once
	OrderService := service2.NewOrderService(repository.NewOrderRepository(db))
	//load balance

	//Hystrix add


	//show watch ip
	common.PrometheusBoot(9092)
	//micro service
	Service := micro.NewService(
		micro.Name("go.micro.service.order"),
		micro.Version("latest"),
		//out service address
		micro.Address(":9085"),
		//registry consul
		micro.Registry(consul),
		//open tracer
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		// add current limiting
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
		//add watch
		micro.WrapHandler(prometheus.NewHandlerWrapper()),


	)


	//service init
	Service.Init()
	//registry handler
	if err := order.RegisterOrderHandler(Service.Server(), &handler.Order{OrderService: OrderService});err!=nil{
		log.Error(err)
	}
	//run service
	if err :=Service.Run();err!=nil{
		log.Error(err)
	}

}