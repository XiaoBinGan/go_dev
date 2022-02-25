package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"user/domain/repository"
	service2 "user/domain/service"
	"user/handler"
	user "user/proto/user"

)

func main() {
	// Create service attr
	srv := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
	)
	//init server
	srv.Init()
	//创建数据库连接
	db, err := gorm.Open("mysql", "root:@/micro?charset=utf8&parseTime=true&loc=Local")
	if err!=nil{
		fmt.Println(err)
	}
	defer  db.Close()

	db.SingularTable(true)
    //只执行一次  数据表初始化
	rp := repository.NewUserRepository(db)
	rp.InitTable()
	//创建服务实例
	userDataService := service2.NewUserDataService(repository.NewUserRepository(db))
	err = user.RegisterUserHandler(srv.Server(), &handler.User{UserdDataService: userDataService})
	if err!=nil{
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
