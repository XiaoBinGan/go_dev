package main

import (
	"fmt"
)
// defer 
//defer语句延迟调用的特性,所以defer语句非常方便的处理资源释放问题,资源清理/文件关闭/解锁及记录时间等
func deferDemo()  {
	fmt.Println("start")
	defer fmt.Println("11111") //先被defer的语句最后被执行，最后被defer的语句，最先被执行
	defer fmt.Println("22222")
	defer fmt.Println("33333")
	fmt.Println("end")
}
func main() {
	deferDemo()
}