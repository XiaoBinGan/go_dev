package main

import (
	"fmt"
)
//递归
//递归适合处理那些问题相同|问题越来越小的场景
//递归一定要有一个明确的退出条件
//永远不要高估自己



// 3! = 3*2*1    
// 4! = 4*3*2*1  
// 5! = 5*4*3*2*1

//计算的n的阶乘(1一直乘到n的值)
func f(n uint64) uint64 {
	if n<1 {
		return 1
	}
	return n*f(n-1)
}

func f1(n int) int{
	if n<1 {
		return 1
	}
	return n*f1(n-1)
}

func main() {
	res  :=f1(9)
	fmt.Println(res)
}