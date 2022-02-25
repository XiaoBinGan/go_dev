package main

import (
	"fmt"
)

//只写
func send(a chan<- int,exitchan chan struct{})  {
	for i := 0; i < 10; i++ {
		a <- i
	}
	close(a)
	var d struct{}
	exitchan <- d
}
//只读
func receive(b<- chan int, exitchan chan struct{}){
	for {
		c,ok :=<- b
		if !ok {
			fmt.Println("error or not enough")
			break
		}
		fmt.Println(c)
	}
	var d struct{}
	exitchan <- d
}


func main() {
	var m chan int      //var 变量 chan <- 类型 只允许写入 不允许读取 读取则报错
	m = make(chan int ,10)
	exitchan :=make(chan struct{},2)
	go send(m,exitchan)
	go receive(m,exitchan)
	totla :=0
	for _ = range exitchan {
		totla++
		if totla ==2{
			break
		}
				
	}
} 