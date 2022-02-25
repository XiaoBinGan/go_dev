package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {
		//如果不叫rand.Seed那么随机数每次都一样
		rand.Seed(time.Now().UnixNano())//添加一个int64的数值让每次随机数都不一样
		random1 :=rand.Int()//int类型的随机数
		random2 :=rand.Intn(10) //0=<X<10 左开右和
		fmt.Println(random1,random2)
	}
}