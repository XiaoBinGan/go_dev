package main

import (
	"time"
	"fmt"
)

func write(a chan int)  {
	fmt.Println("kaishi")
	for i := 0; i < 100; i++ {
		a <- i
		fmt.Println(i)
	}
}

func read(b chan int)  {
	fmt.Println("结束")
	for {
		var c int 
		c = <- b
		fmt.Println(c)
		time.Sleep(time.Second * 1)
	}
}


func main() {
	q :=make(chan int,10)      //新建一个chan 类型为int 长度为10
	go write(q)			       //不停的写入
	go read(q)			       //不停的去读取  只有读取管道里的内容新的内容才能被写入
   time.Sleep(time.Second * 10)  //加一个延迟器

} 