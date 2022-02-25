package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //gorm mysql驱动
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"go_dev/day18/e_commerce/category/common"
	"go_dev/day18/e_commerce/category/domain/repository"
	service2 "go_dev/day18/e_commerce/category/domain/service"
	"go_dev/day18/e_commerce/category/handler"
	category "go_dev/day18/e_commerce/category/proto/category"
)

func main()  {
	//config center
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")

	if err!=nil{
		log.Error(err)
	}
	//registry center
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	//new service
	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8082"),
		micro.Registry(consulRegistry),
	)

	//read mysql conf
	mysqlInfo,err := common.GetMysqlFromConsul(consulConfig,"mysql")
	if err!=nil{
		log.Error(err)
	} //db, err := gorm.Open("mysql", "root:@/micro?charset=utf8&parseTime=true&loc=Local")
	//db, err := gorm.Open("mysql", mysqlInfo.User+":@/micro?charset=utf8&parseTime=true&loc=Local")
	db,err := gorm.Open("mysql",mysqlInfo.User+":@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	//db,err := gorm.Open("mysql",mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")

	if err!=nil{
		log.Error(err)
	}
	defer db.Close()
	//forbid repeat table
   db.SingularTable(true)

	//create database table
	categoryRepository := repository.NewCategoryRepository(db)
   //categoryRepository.InitTable()
	//service init
	service.Init()

	//repository
	categoryDataService := service2.NewCategoryDataService(categoryRepository)
	err = category.RegisterCategoryHandler(service.Server(), &handler.Category{CategoryDataService: categoryDataService})
	if err!=nil{
		log.Error(err)
	}
	//Run service
	if err :=service.Run();err!=nil{
		log.Fatal(err)
	}

}

//{
//"host":"127.0.0.1",
//"user":"root",
//"pwd":"",
//"database":"micro",
//"port":"3306"
//}


