package main

import (
	"encoding/json"
	"fmt"
)
 
//St struct
type St struct {	//这里序列化的时候可以转换成多种文件格式的类型
	Name string `json:"studen_name" db:"name" ini:"names"`  //json.Marshal会将参数Name输出成student_name  
	Age int 	`json:"studen_age"`	  //如果参数Name的首字母小写则会造成参数外部无法访问成了局部变量
}

func main() {
	a :=St{
		Name:"st",
		Age:12,
	}
	//序列化
	data, err :=json.Marshal(a)
	if err !=nil {
 		fmt.Println("json encode a failed err:",err)
		return
	}
	fmt.Println(string(data))
	//反序列化
	str :=`{"studen_name":"张三","studen_name":19}` //这里的参数名优先匹配json别名
	var p St  //接收解析出来的数据
	json.Unmarshal([]byte(str),&p)  //传指针是为了json.UnMarshal内部修改p的值
	fmt.Printf("%#v\n",p)

}