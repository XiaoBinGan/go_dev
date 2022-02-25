package main

import (
	"fmt"
	"context"
	"time"
)
// 本示例通过上下文与一个任意的最后期限，告诉阻塞函数，它应尽快放弃工作，因为它得到它
func main() {
	d :=time.Now().Add(50 * time.Millisecond)//当前时间超过50毫秒超时
	ctx,cancel :=context.WithDeadline(context.Background(),d)//
	defer cancel()
	select {
		case <-time.After(1 * time.Second)://1秒超时报错读取内容
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
	}
} 