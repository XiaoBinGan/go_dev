package main

import (
	"fmt"
)

//空接口

//interface :关键字
//interface{} ;空接口

//空接口
func fn(a interface{}) {
	fmt.Printf("a的类型是%T,a的值是%v\n", a, a)
}

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "alex"
	m1["age"] = 9000
	m1["hobby"] = [...]string{"唱歌", "下棋", "吃好吃的"}
	fmt.Println(m1)

	fn(1)
	fn("string")
	fn(m1)
}
