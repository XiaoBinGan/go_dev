package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn,err :=redis.Dial("tcp","localhost:6379")
	if err!=nil {
		fmt.Println("conn redis failed,err:",err)
		return
	}
	defer conn.Close()
	n,err := conn.Do("lpush","book_list", "abc", "ceg",300,500)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
	for {
		r,err :=redis.String(conn.Do("lpop","book_list"))
		if err!=nil {
			fmt.Println("get adc filed,err:",err)
			return
		}
		fmt.Println(r)
	}







}