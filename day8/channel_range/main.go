package main

import (
	"fmt"
)



func main() {
	a :=make(chan int, 10)
	for i := 0; i < 10; i++ {
		a <- i
	}	
	close(a)	      //结束不关闭的话下面读取很容易造成死锁
	for v := range a {// range 可以直接遍历chan
		_=v
		fmt.Println(v)
	}
}