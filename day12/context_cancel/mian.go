package main

import (
	"time"
	"fmt"
	"context"
)




func gen(ctx context.Context) <-chan int  {//返回一个chan
	dst :=make(chan int)
	n :=1
	go func ()  {//起一个goroute不断地去读取和写入内容到dst当中
		for {//死循环遍历
			select {//防止阻塞
				case <-ctx.Done()://超时或者返回异常或者关闭cancel走到这里
					fmt.Println("i exit")
					return
				case dst <-n://没有返回就持续写入
					n++
			}
		}
	}()
	return dst//返回一个chan合集
}

func test()  {
	ctx,cancel :=context.WithCancel(context.Background())//创建一个超时时间
	defer cancel()//结束的时候关闭上下文
	initChanl :=gen(ctx)//获取写入内容
	for n := range initChanl{//
		fmt.Println(n)
		if n ==5 {//这里返回出发ctx的chanl的返回 会触发ctx.Done关闭
			break
		}
	}
}


func main() {
	test()
	time.Sleep(time.Hour)//
}