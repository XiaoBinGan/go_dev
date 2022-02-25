package main

import (
	"fmt"
)




func main() {
	ch :=make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	
    close(ch)// 不关闭 会导致  致命错误:所有goroutines睡着-死锁!
	for {
		b,ok := <-ch
		if !ok {
			fmt.Println("this param is over")
			return
		}
		fmt.Println(b)
	}



}