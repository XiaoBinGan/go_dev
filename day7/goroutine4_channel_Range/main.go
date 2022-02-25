package main

import (
	"sync"
	"fmt"
)



var wg sync.WaitGroup
//使用for range遍历通道，当通道被关闭的时候就会退出for range。
func chanRange()  {
	var a string
	ch1 := make(chan int)
	ch2 := make(chan int)
		// 开启goroutine将0~100的数发送到ch1中
	go func(ch1 chan int) {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			ch1 <-i
		}
		close(ch1)
		fmt.Scanln(&a)		
		fmt.Println(a)
	}(ch1)
		// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func(ch2 chan int) {
		defer wg.Done()
		for{
			i,ok :=<-ch1
			if !ok{
				break
			}
			ch2<-i*i
		}
		close(ch2)
		}(ch2)
		wg.Add(2)
		// 在主goroutine中从ch2中接收值打印
		for v := range ch2 {
			fmt.Println(v)
		}
	wg.Wait()
}
// 限制通道在函数中只能发送或只能接收。
func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	chanRange()

	// ch1 := make(chan int)
	// ch2 := make(chan int)
	// go counter(ch1)
	// go squarer(ch2, ch1)
	// printer(ch2)
}


// for range从通道循环取值
// 当向通道中发送完数据时，我们可以通过close函数来关闭通道。

// 当通道被关闭时，再往该通道发送值会引发panic，
// 从该通道取值的操作会先取完通道中的值，再然后取到的值一直都是对应类型的零值。