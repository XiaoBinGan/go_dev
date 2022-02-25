package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	//创建连接
	// Conn是一种通用的面向流的网络连接
	// 多个goroutines可以同时调用一个Conn上的方法
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("Error dailing,err:", err.Error())
	}
	//函数返回的时候关闭连接
	defer conn.Close()

	//创建请求头内容
	msg := "Get / HTTP/1.1\r\n"
	msg += "Host:www.baidu.com\r\n"
	msg += "Conncetion:kepp-alive\r\n"
	msg += "User-Agent:Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36\r\n"
	msg += "\r\n\r\n"

	// 创建文件缓存
	n, err := io.WriteString(conn,msg)
	if err!= nil {
		fmt.Println("write string failed,err:", err)
		return
	}
	fmt.Println("send to baidu.com bytes:", n)
	//创建接收的字节数组
	buf := make([]byte, 5000)
	for {
		count, err := conn.Read(buf)
		fmt.Println("count:", count, "err:", err)
		if err != nil {
			break
		}
		fmt.Println(string(buf[0:]))
	}

}
