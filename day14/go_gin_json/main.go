package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
/**
	创建默认的模版引擎
    H is a shortcut for map[string]interface{}
	type H map[string]interface{}
    当客户端以GET/POST方法请求/hello路径时，会执行后面的匿名函数
 */
func main() {
	router:=gin.Default()
	router.POST("/hello", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"massage":"Hellow world!",
		})
	})
	type msg struct{
		Name string	`json:"name"`
		Age int		`json:"age"`
		Message string`json:"message"`
	}
	router.POST("/", func(context *gin.Context) {
	data :=msg{
		Name:"👌",
		Age: 19,
		Message: "success",
		}
		context.JSON(http.StatusOK,data)
	})




	/**
		1.name :=c.Query("string")  get a url params
	      if name == ""  cannot error
		2.name:=c.DefaultQuery("key","defaultKey")
	      if cannot get key ,else get the default value
		3.name, ok:=c.getQuery("key")
	       if !ok{
				cannot get the key
			}

	*/
	router.POST("/query", func(context *gin.Context) {
		//name :=context.Query("name")
		//age :=context.Query("age")
		//name:=context.DefaultQuery("name","浩浩")
		age:=context.DefaultQuery("age","18")
		name,ok:=context.GetQuery("name")
		if !ok {
			fmt.Printf("get name false")
		}
		context.JSON(http.StatusOK,gin.H{
			"name":name,
			"age":age,
		})
	})


	router.Run(":9000",func(c *gin.Context) {
		context.JSON(200,gin.H{
			"massage":"Hellow world!",
		})
	})
}