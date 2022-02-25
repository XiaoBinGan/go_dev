package main

import (
	"io/ioutil"
	"fmt"
	"time"
	"context"
	"net/http"
)


type Result struct {
	r *http.Response   //接收请求的响应
	err error
}

func process()  {
	 ctx,cancel :=context.WithTimeout(context.Background(),time.Second*2)//传入一个上下文 和超时时间返回 
	 defer cancel()//返回的时候直接结束 在第一次调用之后，对CancelFunc的后续调用不执行任何操作。
 	 tr :=&http.Transport{}
	 client :=&http.Client{Transport:tr}
	 c :=make(chan Result, 1)//申明一个chan来存储请求的响应
	 req,err :=http.NewRequest("GET","http://www.baidu.com",nil)//new一个请求的对象
	 if err!=nil{
		 fmt.Println("http request failed,err:",err)
		 return
	 }
	go func ()  {
		 resp,err :=client.Do(req)
		 pack :=Result{r:resp,err:err}
		 c <- pack
	}()
	select {
		case <-ctx.Done()://这个管道如果够读取到数据说明超时了
			tr.CancelRequest(req)//取消http的请求
			res :=<-c//取消之后的报错读取
			fmt.Println("Timeout! err:",res.err)
		case res :=<-c: 
			defer res.r.Body.Close()
			out,_ :=ioutil.ReadAll(res.r.Body)
			fmt.Printf("sever Response:%s",out)
	}
	return
}

func main() {
	process()
}




// Background返回一个非nil的空上下文。它从来没有被取消过，也没有
//值，并且没有期限。它通常被主函数使用，
//初始化、测试和作为传入的顶级上下文
//请求。
// func Background() Context {
// 	return background
// }



// // NewRequestWithContext返回一个给定方法、URL和的新请求
// / /可选的身体。
// //
// //如果提供的主体也是一个io。近,返回
// / /请求。Body被设置为Body并将被客户端关闭
// //方法Do、Post、PostForm和Transport.RoundTrip。
// //
// NewRequestWithContext返回一个适合使用的请求
// / /客户端。做或Transport.RoundTrip。创建使用的请求
// //测试一个服务器处理程序，使用方法中的NewRequest函数
// // net/http/httptest包，使用ReadRequest，或手动更新
// / /请求字段。对于传出的客户机请求，为上下文
// //控制请求及其响应的整个生命周期:
// //获取连接，发送请求，并读取
// //响应标头和正文。请参阅请求类型的文档
// //入站和出站请求字段之间的差异。
// //
// //如果主体的类型是*bytes。缓冲区,*字节。读者,或
// / / *的字符串。阅读器，返回的请求的ContentLength被设置为
// //精确值(而不是-1)，GetBody被填充(因此307和308)
// 重定向可以重放主体)，而主体设置为NoBody，如果
// ContentLength为0。
// func NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*Request, error) {
// 	if method == "" {
// 		// We document that "" means "GET" for Request.Method, and people have
// 		// relied on that from NewRequest, so keep that working.
// 		// We still enforce validMethod for non-empty methods.
// 		method = "GET"
// 	}
// 	if !validMethod(method) {
// 		return nil, fmt.Errorf("net/http: invalid method %q", method)
// 	}
// 	if ctx == nil {
// 		return nil, errors.New("net/http: nil Context")
// 	}
// 	u, err := parseURL(url) // Just url.Parse (url is shadowed for godoc).
// 	if err != nil {
// 		return nil, err
// 	}
// 	rc, ok := body.(io.ReadCloser)
// 	if !ok && body != nil {
// 		rc = ioutil.NopCloser(body)
// 	}
// 	// The host's colon:port should be normalized. See Issue 14836.
// 	u.Host = removeEmptyPort(u.Host)
// 	req := &Request{
// 		ctx:        ctx,
// 		Method:     method,
// 		URL:        u,
// 		Proto:      "HTTP/1.1",
// 		ProtoMajor: 1,
// 		ProtoMinor: 1,
// 		Header:     make(Header),
// 		Body:       rc,
// 		Host:       u.Host,
// 	}
// 	if body != nil {
// 		switch v := body.(type) {
// 		case *bytes.Buffer:
// 			req.ContentLength = int64(v.Len())
// 			buf := v.Bytes()
// 			req.GetBody = func() (io.ReadCloser, error) {
// 				r := bytes.NewReader(buf)
// 				return ioutil.NopCloser(r), nil
// 			}
// 		case *bytes.Reader:
// 			req.ContentLength = int64(v.Len())
// 			snapshot := *v
// 			req.GetBody = func() (io.ReadCloser, error) {
// 				r := snapshot
// 				return ioutil.NopCloser(&r), nil
// 			}
// 		case *strings.Reader:
// 			req.ContentLength = int64(v.Len())
// 			snapshot := *v
// 			req.GetBody = func() (io.ReadCloser, error) {
// 				r := snapshot
// 				return ioutil.NopCloser(&r), nil
// 			}
// 		default:
// 			// This is where we'd set it to -1 (at least
// 			// if body != NoBody) to mean unknown, but
// 			// that broke people during the Go 1.8 testing
// 			// period. People depend on it being 0 I
// 			// guess. Maybe retry later. See Issue 18117.
// 		}
// 		// For client requests, Request.ContentLength of 0
// 		// means either actually 0, or unknown. The only way
// 		// to explicitly say that the ContentLength is zero is
// 		// to set the Body to nil. But turns out too much code
// 		// depends on NewRequest returning a non-nil Body,
// 		// so we use a well-known ReadCloser variable instead
// 		// and have the http package also treat that sentinel
// 		// variable to mean explicitly zero.
// 		if req.GetBody != nil && req.ContentLength == 0 {
// 			req.Body = NoBody
// 			req.GetBody = func() (io.ReadCloser, error) { return NoBody, nil }
// 		}
// 	}

// 	return req, nil
// }
