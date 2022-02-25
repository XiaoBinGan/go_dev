package main

import (
	"time"
	"fmt"
	"runtime"
)

func main() {
	num :=runtime.NumCPU()//NumCPU返回本地机器的逻辑CPU个数
	runtime.GOMAXPROCS(num-1)//GOMAXPROCS设置可同时执行的最大CPU数，并返回先前的设置。 若 n < 1，它就不会更改当前设置。本地机器的逻辑CPU数可通过 NumCPU 查询。本函数在调度程序优化后会去掉。
	for i := 0; i < 16; i++ {
		go func ()  {
			t :=time.NewTicker(time.Second)//NewTicker返回一个新的Ticker，该Ticker包含一个通道字段，并会每隔时间段d就向该通道发送当时的时间。它会调整时间间隔或者丢弃tick信息以适应反应慢的接收者。如果d<=0会panic。关闭该Ticker可以释放相关资源。
			for{
				select {
				case <-t.C:
					fmt.Println("timeout")
				}
				t.Stop()
			}
		}()
	}
	time.Sleep(time.Second * 100)//携程执行的非常快所以设一个定时器才能看清执行结果
	fmt.Println(num)
}