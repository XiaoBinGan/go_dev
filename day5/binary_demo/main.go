package main

import (
	"fmt"
)

const (
	eat   int = 4  //十进制的100就是二进制的4
	sleep int = 2	
	da    int =1
 )
//111二进制位图
//左边的1表示吃饭   100
//中间的1便是睡觉  	010
//右边的1便是打豆豆 001

func f(arg int)  {
	fmt.Printf("%b\n",arg)
}
func main() {
	f(eat | da)	
	f(eat | da | sleep)	
}