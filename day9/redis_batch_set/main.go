package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c,err :=redis.Dial("tcp","localhost:6379")
	if err!=nil{
		fmt.Println("connection failed,err:",err)
		return
	}
	c.Close()
	_,err =c.Do("MSet","abc",100,"efg",300) //批量的去新增
	if err!=nil{
		fmt.Println("Mset failed ,err:",err)
		return
	}

	r,err :=redis.Ints(c.Do("MGet","abc","efg"))
	if err!=nil{
		fmt.Println("get abc failed,err:",err)
		return
	}
	for _, v := range r {
		fmt.Println(v)	
	}



}