package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


/**
	直接取指定URL后面的路径为参数
	切记这里不能用根路径直接取后面的参数
	会导致后面的路径全部失效
	建议使用二级或者三级路径去分开
*/

func main() {
	r:=gin.Default()
	r.GET("/user/:name/:age", func(context *gin.Context) {
		name :=context.Param("name")
		age :=context.Param("age")
		data:=map[string]interface{}{
			"name":name,
			"age":age,
		}
		context.JSON(http.StatusOK,data)
	})
	r.GET("/blog/:year/:month", func(context *gin.Context) {
		year :=context.Param("year")
		month :=context.Param("month")
		context.JSON(http.StatusOK,gin.H{
			"year":year,
			"month":month,
		})
	})

	r.Run(":9090")
}
