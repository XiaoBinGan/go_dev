package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c,err :=redis.Dial("tcp","localhost:6379")
	if err!=nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	defer c.Close()
	//可以设置参数过期时间
	r,err := c.Do("expire", "abc", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)

}