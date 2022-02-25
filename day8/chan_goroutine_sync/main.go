package main

import (
	"fmt"
)



func calc(a chan int,b chan int,c chan bool)  {
	for v := range a {
		flag :=true
		for i := 2; i < v; i++ {
			if v%i ==0 {
				flag=false
				break
			}
		}
		if flag {
			b <- v
		}
	}
	fmt.Println("calc is over exit")
	c <- true
}


func main() {
	a := make(chan int, 1000)
	b := make(chan int, 1000)
	c := make(chan bool, 500)
	go func ()  {	
		for i := 0; i < 1000; i++ {
			a <- i
		}
		close(a)
	}()
	for i := 0; i < 500; i++ {  //启用500个线程去计算
			go calc(a,b,c)
	}

	go func ()  {
		for i := 0; i < 500; i++ {
			<-c
		}
		fmt.Println("exit ok you can  close  b")
		close(b)
	}()

	for v := range b {
		fmt.Println(v)
	}

	

}
