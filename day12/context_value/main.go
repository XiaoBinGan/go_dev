package main

import (
	"fmt"
	"context"
)

func process(ctx context.Context)  {
	res,ok :=ctx.Value("trace_id").(int)//获取的类型的强制转换成int类型防止出错
	if !ok{//判断一下是否读取成功
		res = 1231233123231
	}
	fmt.Printf("ret:%d\n",res)


	s,_ :=ctx.Value("session").(string)
	fmt.Printf("session:%s\n",s)
}



func main() {
	ctx := context.WithValue(context.Background(),"trace_id",12313)//返回一个上下文context
	ctx =context.WithValue(ctx,"session","aqwasdqesad")//用上一个上下文对象
	process(ctx)
}