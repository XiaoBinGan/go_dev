package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



/**
	1.1。创建默认的routergin.Default()
	1.2。读取html文件router.LoadHTMLFiles("./xxxxx1.html","./xxxxx2.html")
	1。3	r.LoadHTMLGlob("这里写的是路径")  r.LoadHTMLFiles("这里加载的是文件")
	2.3。通过router创建login接口 使用匿名函数func(context *gin.Context){
		context.HTML(https.StausOk,"index.html",nil)
	}

	2.1 一个请求对应一个响应所以这里需要在定义一个post接口去接受前台的form的post请求
	2.2 使用context.PostForm("key")来接收前台的form的name为key的value

	router.run()
*/

func main() {
	r :=gin.Default()
	r.LoadHTMLFiles("./index.html","./home.html")
	r.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK,"index.html",nil)
	})

	r.POST("/login", func(context *gin.Context) {
		username :=context.PostForm("username")
		//password :=context.PostForm("password")//即使这里的password没传也不会报错这里会接收空字符串
		//password :=context.DefaultPostForm("password","XXXXXX")//有password这个key就不会取后面的XXXX默认值
		//password :=context.DefaultPostForm("xxxx","1231231313") //这里如果没有xxxx这个参数就会取默认的1231231313为password
		password,ok:=context.GetPostForm("123123")//这里只要取到key ok就是true娶不到就执行下面的判断条件
		if !ok{
			password="xxxxxx"
		}
		//gin.H{}   这里不能忘记的是这里 map[string]interface{}
		context.HTML(http.StatusOK,"home.html",gin.H{
			"username": username,
			"password": password,
		})
	})






	r.Run(":9000")
}