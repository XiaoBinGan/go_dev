package main

import "fmt"

//1
type  addFunc func(int,int) int //golang当中尽量使用驼峰命名
								//使用type注册类型使用的时候申明类型的时候需要是addFunc这个名称

//2
func add (a,b int) int{
	return a+b
}	  

//3
func test(a addFunc,b,c int) int {
	return a(b,c)
}	
//4							
func main()  {
	c :=add
	// fmt.Printf("",&c)
	sum := test(c,12,34)
	fmt.Printf("%d /n",sum)	
}