package main

import (
	"fmt"
)

//函数的类型
func f1()  {
	fmt.Println("string")
}

func f2() int {
	return 10
}     
//函数作为参数的类型
func f3(x func() int)  {
	ref := x()
	fmt.Println(ref)
}
func f4()int{
	fmt.Println("f4")
	return 4
}  
//参数是一个函数  返回值也是一个函数但是函数有参数和没有参数的类型是不一样的
func f5(x func()int) func()int {
	return x
} 

//将函数当成返回值



func main() {
	a :=f1
	fmt.Printf("%T\n" ,a)
	b :=f2
	fmt.Printf("%T\n" ,b)

	f3(f2)
	fmt.Printf("%T\n" ,f3)

	f6 :=f5(f4)

	f6()
	fmt.Printf("%T\n", f6)


}