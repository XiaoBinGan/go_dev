package main

import (
	"runtime"
	"time"
	"fmt"
)





func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}
func main() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(2)//将逻辑核心数设为2，此时两个任务并行执行，默认跑满
	go a()
	go b()
	time.Sleep(time.Second)
}
// Go语言中的操作系统线程和goroutine的关系：

// 一个操作系统线程对应用户态多个goroutine。
// go程序可以同时使用多个操作系统线程。
// goroutine和OS线程是多对多的关系，即m:n。