package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)
 
func fnGet() {
	res, err := http.Get("0.0.0.0:8080/")
	if err != nil {
		fmt.Println("get er:", err)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("get data err:", err)
		return
	}
	fmt.Println(string(data))
}

/*get请求带参数
  创建的一个请求的url
  使用url.Values{}来存储参数
  使用url.RawQuery对象来接收Encode的参数,并用参数
  使用url.String()来将url.RawQuery和url.ParseRequestURi拼接起来
  http.Get(url.String()) 向指定的地址发送get请求url后面拼接参数
  容错防止http请求失败或者超时
  使用ioutil.ReadAll()读取所有的数据返回
  容错防止数据读取失败
  将返回的数据tostring然后操作
*/
func fnGet1() {
	uri := "http://0.0.0.0:8080/get"
	//url param
	data := url.Values{}
	data.Set("name", "Ami")
	data.Set("age", "18")
	// fmt.Println(data)//data是个map所以需要编码成字符串
	u, err := url.ParseRequestURI(uri)
	if err != nil {
		fmt.Printf("parse url requestUrl failed,err:%v\n", err)
		return
	}
	u.RawQuery = data.Encode() //url.String()接收一个u.RawQuery所以将data.Encode将map转换成字符串赋值给u.RawQuery
	fmt.Println(u.String())    //http://0.0.0.0:8080/post?name=Ami
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("Get failed,err:%v\n", err)
		return
	}
	bt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ReadAll failed err:%v\n", err)
		return
	}
	fmt.Println(bt)
	type res struct { //申明一个接受数据的struct
		Name string `json:"name"`
		Age  string `json:"age"`
	}
	var r *res                   //用一个res累心的指针变量来接收服务端返回的数据
	err = json.Unmarshal(bt, &r) //将返回的数据解码放入指定格式的res的地址
	fmt.Printf("data:%+v\n", r)
}

/*发送post请求的客户端
  创建请求地址
  确定请求的数据类型
	  1.表单"application/x-www-form-urlencoded"
	  	数据格式字符串拼接“name=ali&age=13”
	2.json"application/json"
		数据格式json`{"name":"ali","age",13}`
  使用http.Post(Url,contentType,strings.NewReader(data))
  容错
  defer关闭数据响应body的流
  读取返回body的内容
  容错
*/
func fnPost() {
	uri := "http://0.0.0.0:8080/post"
	//application/x-www-form-urlencoded
	// contentType := "application/x-www-form-urlencoded"
	// data := "name=ali&age=13"
	contentType :="application/json"
	data :=`{"name":"ali","age":13}`
	res, err := http.Post(uri, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("request failed err:", err)
		return
	}
	defer res.Body.Close()
	bt, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read body failed ,err:", err)
		return
	}
	fmt.Println(string(bt))
}

func main() {
	// fnGet()
	// fnGet1()
	fnPost()
}
