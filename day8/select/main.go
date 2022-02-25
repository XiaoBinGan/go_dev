package main

import (
	"fmt"
	"time"
)


func main() {
	var a chan int
	var b chan int
	a = make(chan int, 10)
	b = make(chan int, 10)

	go func ()  {
		for i := 0; i < 1000; i++ {
			a <- i
			time.Sleep(time.Second *1)
			b<- i*i
			time.Sleep(time.Second *1)
		}
	}()

	for {
		select {
			case c :=<-a :
				fmt.Println(c)
			case v :=<-b :
				fmt.Println(v)
			default:
				fmt.Println("time out")
				time.Sleep(time.Second)
		}
	}


	
}