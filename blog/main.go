package main

import (
	"github.com/gin-gonic/gin"
	"go_dev/blog/controller"
	db "go_dev/blog/dao/db"
)

func init(){
	//parseTime=true 将MySQL中的时间类型自动解析为go语言结构体中的时间类型
	//dns :="root:@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	var dns ="root:@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := db.Init("mysql", dns)
	if err!=nil{
		panic(err)
	}
}

func main() {
	r :=gin.Default()
	r.POST("/",controller.CategoryHandler)
	r.Run(":8080")
}
