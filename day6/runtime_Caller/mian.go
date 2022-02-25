package main

import (
	"path"
	"fmt"
	"runtime"
)
func f()  {
	pc,file,line,ok :=runtime.Caller(0)//0代表申明时候行数1表示上一层2表示在上一次
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName :=runtime.FuncForPC(pc).Name()//拿到执行的方法名
	fmt.Println(funcName)
	fmt.Println(file)//文件名称全路径
	fmt.Println(path.Base(file))//获取到路径最后的文件名称
	fmt.Println(line)//拿到执行runtime.Caller函数的函数
}
func f1()  {
	f()
}
func main() {
	f1()
}