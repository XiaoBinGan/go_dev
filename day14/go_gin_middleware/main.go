package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
/**
	全局中间件
	@return type HandlerFunc func(*Context)
	1。中间件请求前执行内容
	2。c.Next()处理后续的中间件 没有就处理请求中的函数
	3。等待中间件后续的函数执行完成在继续执行这里的中间件的内容
	4.中间件全部执行完成结束返回给前端
*/
func middleware()gin.HandlerFunc  {
	return func(c *gin.Context) {
		start :=time.Now()
		c.Next()
		since :=time.Since(start)
		time.Sleep(time.Second*6)
		fmt.Printf("程序耗时：%v\n",since)
	}
}
/**
	这里省略常规的路由注册和服务启动的注释
	1.中间件注册router.Use(middleware())
	2.创建请求的分组rel:=router.Group("reltivePath")
	3.为了代码规范使用中括号
	{这里写具体的请求这里的path是基于rel的二级路径
		rel.Get("path",handlerFunc)
	}
*/
func main() {
	r:=gin.Default()
	r.Use(middleware())
	user :=r.Group("user")
	{
		user.GET("/",userlogin)
		user.GET("/d",userlg)
	}
	r.Run(":9000")
}
func userlogin(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"name":"jack",
	})
}
func userlg(c *gin.Context)  {
	time.Sleep(time.Second*2)
}