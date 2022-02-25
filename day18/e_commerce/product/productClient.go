package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"go_dev/day18/e_commerce/product/common"
	go_micro_service_product "go_dev/day18/e_commerce/product/proto/product"
)

func main()  {
	//config center
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	//tracing
	tracer, io, err := common.NewTracer("go.micro.service.product.client", "localhost:6831")
	if err!=nil{
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(tracer)

	service := micro.NewService(
		micro.Name("go.micro.service.product.client"),
		micro.Version("latest"),
		micro.Address("1270.0.0.1"),
		//add register center
		micro.Registry(consul),
		//bind tracer
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
	)
	productService := go_micro_service_product.NewProductService("go.micro.service.product", service.Client())
	productInfo := &go_micro_service_product.ProductInfo{
		ProductName:        "iphone16",
		ProductSku:         "iphone",
		ProductPrice:       5666,
		ProductDescription: "最新苹果手机",
		ProductCategoryId:  1,
		ProductImage: []*go_micro_service_product.ProductImage{
			{
				ImageName: "iphone16",
				ImageCode: "iphone01",
				ImageUrl:  "https://www.apple.com.cn/v/iphone/home/av/images/overview/compare/compare_iphone_11__bzjboswm5hbm_small_2x.jpg",
			},
			{
				ImageName: "iphone17",
				ImageCode: "iphone02",
				ImageUrl:  "https://www.apple.com.cn/iphone/home/images/overview/hero/iphone_12__d51ddqcc7oqe_small_2x.jpg",
			},
		},
		ProductSize: []*go_micro_service_product.ProductSize{
			{
				SizeName: "iPhone1",
				SizeCode: "01",
			},
			{
				SizeName: "iPhone2",
				SizeCode: "02",
			},
		},
		ProductSeo: &go_micro_service_product.ProductSeo{
			SeoTitle:       "iPhone",
			SeoKeywords:    "new  iphone",
			SeoDescription: "最新iPhone18",
			SeoCode:        "18",
		},
	}
	responseProduct, err := productService.AddProduct(context.TODO(), productInfo)
   if err!=nil{
   			fmt.Println(err)
   }
   fmt.Println(responseProduct)
}