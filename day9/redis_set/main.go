package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	// 创建redis链接 Dial最终返回一个Dial
	c,err :=redis.Dial("tcp","localhost:6379")
	if err!=nil {
		fmt.Println("Connect tp redis error:",err)
		return 
	}
	//函数返回的时候关闭链接
	defer c.Close() 
	n,err :=c.Do("Set","abc",100)
	fmt.Println(n)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	r,err :=redis.Int(c.Do("Get","abc"))
	if err!=nil{
		fmt.Println("get abc failed,err:",err.Error())
		return
	}
	fmt.Println(r)

}