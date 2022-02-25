package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
	1.HTTP重定向
		直接将请求转发到另外的内部或者外部地址
		context.Redirect(http.StatusMovedPermanently,"https:XXXX.com")
		将请求永久重定向到后面的地址 permanently永久地
	2.路由重定向
		将a的请求转发到b的接口上让b去处理这次请求和响应
		context.Request.URL.Path="转发的接口地址"
		router.HandleContext(context) 将·上下文·传入enginen内部的处理函数 供下面的函数调用

*/



func main() {
	r:=gin.Default()
	r.GET("/test", func(context *gin.Context) {
			context.Redirect(http.StatusMovedPermanently,"https://www.sogou.com")
	})

	r.GET("/a", func(context *gin.Context) {
			context.Request.URL.Path="/b"
			r.HandleContext(context)
	})

	r.GET("/b", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"message":"ok",
		})
	})

	r.Run(":9000")
}
