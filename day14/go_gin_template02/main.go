
package main

import (
"fmt"
"html/template"
"net/http"
)

func hd(w http.ResponseWriter,r  *http.Request)  {
	/*
		自定义函数k
		要么只有一个返回值，要么有两个返回值，第二个返回值必须是error类型
	*/
	k := func(name string)(string,error) {
		return  name +"喜欢打篮球", nil
	}
	//自定义模版   下面两个文件名称需要对应
	t:=template.New("index.tpl")
	//告诉模版引擎，现在是多了一个自定义函数 对应模版中的key返回值
	t.Funcs(template.FuncMap{
		"hobby":k,
	})
	_,err:=t.ParseFiles("index.tpl")
	if err!=nil{
		fmt.Printf("pars template failed ,err%v\n",err)
		return
	}
	name :="浩浩"
	t.Execute(w,name)



}



func main() {
	http.HandleFunc("/",hd)
	http.ListenAndServe(":9000",nil)
}