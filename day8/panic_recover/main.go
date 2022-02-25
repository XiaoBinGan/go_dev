package main

import (
	"time"
	"fmt"
)
//注意的事recover必须搭配defer使用
//defer一定要在可能发生panic的语句之前定义
func test()  {
	defer func ()  {//函数返回的时候执行defer
		  // 发生宕机时，获取panic传递的上下文并打印
		err :=recover();
		if err !=nil{//这里判断是不是能够恢复如果不能就直接报错可以则继续尝试恢复一下
			fmt.Println("painc:",err)
		}
	}()	
	 var m map[string]int
	m["stu"] = 100
}

func calc() {
	for {
		fmt.Println("i'm calc")
		time.Sleep(time.Second)
	}
}

func main() {
	// a :=map[string]int{
	// 	"age":123,
	// 	"score":100,
	// }
	// fmt.Println(a)

	go test()
	for i := 0; i < 2; i++ {
		go calc()
	}
	time.Sleep(time.Second)
}