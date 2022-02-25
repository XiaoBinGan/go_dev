package main

import (
	"fmt"
	"time"
	"net"
	"net/http"
)


var url = []string{
	"https://www.baidu.com",
	"https://google.com",
	"https://taobao.com",
}


func main() {
	for _,v :=range url {
		c :=http.Client{
			
			Transport: &http.Transport{ 
				Dial:func(network, addr string) (net.Conn, error){
					timeout :=time.Second*2//设置两秒超时
					return net.DialTimeout(network,addr,timeout)
				},
			},
		}

		resp,err :=c.Head(v)
		if err!=nil{
			fmt.Println("ger head is failed err:",err)
		}
		fmt.Println(resp)
	}
}

// Transport是支持HTTP、HTTPS和HTTP代理(用于HTTP或带CONNECT的HTTPS)的RoundTripper实现。
// 默认情况下，传输缓存连接以供将来重用。这可能会在访问许多主机时留下许多打开的连接。
// 可以使用传输的CloseIdleConnections方法和MaxIdleConnsPerHost和DisableKeepAlives字段来管理此行为。
// 传输应该被重用，而不是根据需要创建。传输对于由多个goroutines并发使用是安全的。
// 传输是用于发出HTTP和HTTPS请求的低级原语。有关高级功能，如cookie和重定向，请参阅Client。
// 传输对HTTP url使用HTTP/1.1，对HTTPS url使用HTTP/1.1或HTTP/2，这取决于服务器是否支持HTTP/2以及传输的配置方式。
// DefaultTransport支持HTTP/2。要显式地在传输上启用HTTP/2，请使用golang.org/x/net/http2并调用ConfigureTransport。有关HTTP/2的更多信息，请参阅包文档。
// 状态码在1xx范围的响应要么自动处理(100 expect-continue)，要么忽略。
// 一个例外是HTTP状态码101(交换协议)，它被认为是一种终端状态，并通过往返返回。
// 要查看被忽略的1xx响应，请使用httptrace跟踪包的ClientTrace.Got1xxResponse。
// 传输只有在遇到网络错误时才重试请求，如果请求是幂等的，或者没有主体，或者有它的请求。
// GetBody定义。如果HTTP请求具有HTTP方法GET、HEAD、OPTIONS或TRACE，则认为它们是幂等的;或者它们的头映射包含一个“等势键”或“x -等势键”条目。
// 如果等幂键值是一个零长度的片，则请求被视为等幂，但不发送消息头。