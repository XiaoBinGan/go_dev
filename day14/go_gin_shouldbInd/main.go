package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
context.Shouldbind
*/


/**
	这里申明一个结构体
	form tag    form表单提交
	json tag	post提交json数据格式
	binding:"required"  必传字段
	分别用来告诉 context.ShouldBind()什么请求类型的参数对应结构体的什么名称
	ShouldBind()方法使用反射来对应结构体内部参数和请求参数，所以请注意结构体内部的参数尽量大写开头
	ShouldBind可以对应任何请求类型的参数只要在结构体的参数后面加上相对于的tag
*/
type Student struct {
	UserName string `form:"username" json:"username" binding:"required" `
	PassWord string `form:"password" json:"password" binding:"required" `
}


func main() {
	r:=gin.Default()
	r.POST("/form", func(context *gin.Context) {
		var u Student
		err :=context.ShouldBind(&u)
		if err!=nil{
			context.JSON(http.StatusOK,gin.H{
				"errorMassage":"400",
			})
		}else {
			fmt.Printf("%+#v\n",u)
			context.JSON(http.StatusOK,gin.H{
				"massage":"ok",
			})
		}
	})
	r.POST("/json", func(context *gin.Context) {
		var u Student
		err :=context.ShouldBind(&u)
		if err!=nil{
			context.JSON(http.StatusOK,gin.H{
				"errorMassage":"400",
			})
		}else {
			fmt.Printf("%+#v\n",u)
			context.JSON(http.StatusOK,gin.H{
				"massage":"ok",
			})
		}
	})
	r.GET("/get", func(context *gin.Context) {
		var u Student
		err :=context.ShouldBind(&u)
		if err!=nil{
			context.JSON(http.StatusOK,gin.H{
				"errorMassage":"400",
			})
		}else {
			fmt.Printf("%+#v\n",u)
			context.JSON(http.StatusOK,gin.H{
				"massage":"ok",
			})
		}
	})
	r.Run(":9000")

}
