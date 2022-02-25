package main

import (
	"time"
	"fmt"
	"sync"
)

func main() {
	 wg :=sync.WaitGroup{}
	 var count int =10
	 for i := 0; i < count; i++ {
		 wg.Add(1)//使用一次+1
		 go calc(&wg,i)//将 
	 }
	 wg.Wait()//当数值为0的时候直接退出
	 fmt.Println("all goroute finis")
}
//接收类型为sync.WaitGroup的变量地址 因为这里是指针 
func calc(w *sync.WaitGroup,i int){
	fmt.Println("calc:",i)
	time.Sleep(time.Second)
	w.Done()//使用一次-1
}