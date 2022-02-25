package main

import (
	"log"
	"fmt"
	"io"
	"net/http"
)



const form = `<html><body><form action="#" method="post" name="bar">
                    <input type="text" name="in"/>
                    <input type="text" name="in"/>
                     <input type="submit" value="Submit"/>
			 </form></body></html>`
			 
func SServer(w http.ResponseWriter,request *http.Request )  {
	io.WriteString(w, "hello world")
	panic("err err")
}

func FormServer(w http.ResponseWriter,request *http.Request)  {
	w.Header().Set("Content-Type","text/html")
	switch request.Method {
	case "GET":
		io.WriteString(w,form)    //将文件写到respones里返回给客户端
 	case "POST":
 		request.ParseForm()		  //解码form参数
 		io.WriteString(w,request.Form["in"][0]) //获取表单中的name为in的第0位
 		io.WriteString(w,"\n")
 		io.WriteString(w,request.Form["in"][1])
	}
}
func logPanics(handle http.HandlerFunc) http.HandlerFunc { //处理异常的panic让程序能够正常的运行
	return func (writer http.ResponseWriter,req *http.Request)  {//将要处理的响应和请求进行recover操作防止程序panic
		defer func(){//函数执行结束的时候触发defer检查函数执行是否有异常
			if x :=recover();
			x !=nil{
				log.Printf("[%v] caught panic :%v",req.RemoteAddr,x)
			}
		}()
		handle(writer,req)//没问题将传入的函数正常执行
	}
}
func main() {
	http.HandleFunc("/s",logPanics(SServer))//HandleFunc在DefaultServeMux中为给定的模式注册处理程序函数。ServeMux文档解释了模式是如何匹配的。
	http.HandleFunc("/f",logPanics(FormServer))
	err :=http.ListenAndServe("0.0.0.0:8080",nil)
	if err!=nil{
		fmt.Println("server err,err:",err)
	}
}

// RemoteAddr允许HTTP服务器和其他软件进行记录
// 发送请求的网络地址，通常为
// 日志记录。该字段不是由ReadRequest和填写的
// 没有定义格式。此包中的HTTP服务器
// 在调用前将RemoteAddr设置为一个“IP:port”地址
// 处理程序。
// HTTP客户端会忽略此字段。