package main

import (
	"github.com/gin-gonic/gin"

	"gorm_connect/conf"
	"gorm_connect/handler"
	"gorm_connect/utils"
)

func init() {
	conf.Init()  //init config
	utils.Init() //init database

}

func main() {
	Route := gin.Default()
	v1 := Route.Group("/v1/")
	{
		v1.POST("/adduser", handler.AddUser)
		v1.POST("/login", handler.Login)
	}
	Route.Run(":8080")
}
