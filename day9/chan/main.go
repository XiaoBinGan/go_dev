package main

import (
	"time"
	"fmt"
)

func main() {
	a :=make(chan int, 10)
	 go func ()  {
		 var i int
			for{
				select {//防止线程夯住
				case a<-i://循环写入
				default:
					fmt.Println("channel is full")
					time.Sleep(time.Second)//怕执行的太快看不见
				}
				i++
			}
			 
	 }()
	 for{//循环读取写入的值
		 v :=<-a
		 fmt.Println(v)
	 }
}