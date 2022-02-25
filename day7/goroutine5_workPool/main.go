package main

import (
	"time"
	"fmt"
)


//指定工作的线程数 ,只获取任务,只读取任务
func worker(n int,jobs <-chan int ,result chan<- int)  {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n",n,j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n",n,j)
		result <- j*2
	}
	
}

func main() {
	jobs :=make(chan int,50)	
	result :=make(chan int,50)	
	for i := 1; i <=3; i++ {//创建三个线程去读取和写入资源
		go worker(i,jobs,result)//指定工作的线程数 输出chan和写入chan
	}
	for i := 1; i <=5; i++ {//写入资源
		jobs <-i
	}
	for a := 0; a < 5; a++ {//读取资源
		<-result
	}

}