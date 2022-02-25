package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
Hello name
params w htpp.ResponesWriter
params r *http.Request
*/
func Hello(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("handle hello")
	fmt.Println(r)
	fmt.Fprintf(w,"hello")//将数据写入w返回给客户端
}
/*接收Url后面带参数的get请求
  因为读取接接收的参数也是在对流进行操作所以也是需要关闭的即使是get请求也别忘记关闭r.body.close()
  使用http中的request对象的url对象的Query()方法来获取url后面的参数
  返回一个type Values map[string][]string Valeuls类型
  使用Valus的构造方法Get(key)获取Values的值
*/
func gt(w http.ResponseWriter,r *http.Request )  {
	type res struct{
		Name string	`json:"name"`
		Age  string	`json:"age"`
	}
	defer r.Body.Close()
	data :=r.URL.Query()
	name :=data.Get("name")
	age :=data.Get("age")
	rp :=&res{
		Name:name,
		Age:age,
	}
	jsbyte,err :=json.Marshal(rp)
	fmt.Println(jsbyte)
	if err!=nil{
		fmt.Println("json encode failed,err:", err)
		return
	}
	fmt.Println(name,age)
	w.Write(jsbyte)	
}
/*接收Post请求
  结束的时候关闭流
  1.请求类型是application/x-www-form-urlencoded时解析form数据
	  r.ParseForm()反序列话form数据
	  r.PostForm.Get("key")来获取客户端传入的数据
	  
  2.请求类型是application/json时从r.Body读取数据
	b, err := ioutil.ReadAll(r.Body)
	容错
  返回客户端数据
*/
func pt(w http.ResponseWriter,r *http.Request) {
	defer r.Body.Close()
	r.ParseForm()
	fmt.Println(r.ParseForm())
	fmt.Println("form",r.PostForm.Get("name"),r.PostForm.Get(("age")))
	bt,err :=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Read body failed err:",err)
		return
	}
	fmt.Println("json",string(bt))
	var answer=`{"code":"ok"}`
	w.Write([]byte(answer)) 
}




func main() {
	http.HandleFunc("/",Hello)
	http.HandleFunc("/get",gt)
	http.HandleFunc("/post",pt)
	err :=http.ListenAndServe("0.0.0.0:8080",nil)
	if err!=nil{
		fmt.Println("http listen filed")
		return
	}	
}