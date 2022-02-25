package main

import "fmt"

func add(a,b int)(c int){ //返回类型定义成具体参数名+类型的时候 return 可以为空
	 c =a +b
	 return
}

func rud(a,b int)(e ,f int){ //返回类型定义成具体参数名+类型的时候 return 可以为空
	e =a+b
	f =b*a
	return
}

func main() {
	fmt.Printf("%d \n",add(12,34))	
	
	
	
	e,_ :=rud(12,23) //暂时用不上的值可以使用_暂时
	fmt.Printf("%d\n",e)	
	
}