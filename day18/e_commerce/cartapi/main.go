package main

import (
	"context"
	"fmt"
	go_micro_service_cart "github.com/XiaoBinGan/cart/proto/cart"
	"github.com/XiaoBinGan/common"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"go_dev/day18/e_commerce/cartapi/handler"
	"go_dev/day18/e_commerce/cartapi/proto/cartApi"
	"net"
	"net/http"
)

func main()  {
	//registry center
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	//add jaeger tracer
	tracer,io ,err:=common.NewTracer("go.micro.api.cartApi","localhost:6831")
	if err!=nil{
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(tracer)
    //add protector
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()

	go func() {
		//监控并且上报状态
		if err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", "9096"), hystrixStreamHandler);err!=nil{
			log.Error(err)
		}
	}()


	service := micro.NewService(
		micro.Name("go.micro.cartApi"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8086"),
		//add registry center
		micro.Registry(consul),
		//add open trace new client wrapper
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		//add hystrix
		micro.WrapClient(NewClientHystrix()),
		//add load balance
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)

   //initialise service
	service.Init()

	//new cart service
	cartService := go_micro_service_cart.NewCartService("go.micro.service.cart", service.Client())
	//Registry Handler server
   if	err := cartApi.RegisterCartApiHandler(service.Server(), &handler.CartApi{CartService: cartService});err!=nil{
	   log.Error(err)
   }


	if err := service.Run();err!=nil{
		log.Error(err)
	}
}


type clientWrapper struct {
	client.Client
}
//Do runs your function in a synchronous manner, blocking until either your function succeed
//or an error is returned, including hystrix circuit errors
//func Do(name string, run runFunc, fallback fallbackFunc) error{}
func(c  *clientWrapper)Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error{
  return hystrix.Do(req.Service()+"."+req.Endpoint(),
  	func()error{
  	  //run normal action
  		fmt.Println(req.Service()+"."+req.Endpoint())
  		return c.Client.Call(ctx,req,rsp,opts...)
	},//failed action
	func(err error)error{
        fmt.Println(err)
        return err
	})
}

func NewClientHystrix()client.Wrapper  {
	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}
}
//--registry=consul --registry_address=192.168.8.108:8500 api --handler=api