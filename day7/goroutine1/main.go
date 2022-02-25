package main

import (
	"fmt"
	"sync"
)







var wg sync.WaitGroup //来实现goroutine的同步
// wg.Add(1)//启动一个goroute就会+1
// wg.Done() //goroutine结束就登记-1
// wg.Wait() // 等待所有登记的goroutine都结束
func hellow(i int)  {
	defer wg.Done() //goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)//启动一个goroute就会+1
		go hellow(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}

//一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数