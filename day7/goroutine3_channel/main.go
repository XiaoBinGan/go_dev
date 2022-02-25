package main

import (
	"fmt"
)



func demo1(){//这种方式必须创建缓冲区不然会报错死锁
	a :=make(chan interface{},10)
	a<-1
	fmt.Printf("type:%T,value:%+v\n",a,a)
}

func recv(c chan int) {//无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道。
	ret := <-c
	fmt.Println("接收成功", ret)
}
func send()  {
	ch := make(chan int)
	go recv(ch) //B) 启用goroutine从通道接收值  
	ch <- 10//A)这里先发送的话会阻塞 直到另一个goroutine接收,发送才能成功
	fmt.Println("发送成功")
	// 无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，
	// 两个goroutine将继续执行。相反，如果接收操作先执行，接收方的goroutine将阻塞，
	// 直到另一个goroutine在该通道上发送一个值。
	// 使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道。
}

func main() {
	// demo1()	
	send()

}
//在程序启动时，Go程序就会为main()函数创建一个默认的goroutine。
// 当main()函数返回的时候该goroutine就结束了，
// 所有在main()函数中启动的goroutine会一同结束，
// main函数所在的goroutine就像是权利的游戏中的夜王，
// 其他的goroutine都是异鬼，夜王一死它转化的那些异鬼也就全部GG了。










//1.channel类型
// channel是一种类型，一种引用类型。声明通道类型的格式如下：

// var 变量 chan 元素类型
// var ch1 chan int   // 声明一个传递整型的通道
// var ch2 chan bool  // 声明一个传递布尔型的通道
// var ch3 chan []int // 声明一个传递int切片的通道
// var ch3 chan interface{} // 声明一个传递interface的通道
//2.创建channel
// 通道是引用类型，通道类型的空值是nil。

// var ch chan int
// fmt.Println(ch) // <nil>
//3.声明的通道后需要使用make函数初始化之后才能使用。

// 创建channel的格式如下：

// make(chan 元素类型, [缓冲大小])