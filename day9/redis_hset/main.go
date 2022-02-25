package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c,err :=redis.Dial("tcp","localhost:6379")
	if err!= nil {
		fmt.Println("redis connection failed ,err:",err)
		return
	}
	defer c.Close()

	n,err :=c.Do("HSet","book","css世界","golang高级编程")
	if err!=nil {
		fmt.Println("HSet failed ,err:",err)
		return
	}
	fmt.Println(n)
	r,err:= redis.String(c.Do("HGet","book","css世界"))
	if err!=nil {
		fmt.Println("redis string failed,err:",err)
		return
	}
	fmt.Println(r)
}