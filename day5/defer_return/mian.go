package main

import (
	"fmt"
)


// 在Go语言的函数中return语句在底层并不是原子操作，
// 它分为给返回值赋值和RET指令两步。
// 而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。


func f1() int {
	x := 5				//第二步	函数返回时先赋值
	defer func() {	
		x++				//第三步	这时候修改的是x不再是返回值了
	}()		
	return x			//第-步	将赋值的参数返回	x=5	
}

func f2() (x int) {
	defer func() {
		x++				//第二步 这里又将返回值X进行了添加操作
	}()
	return 5			//第一步 因为返回值提前啊定义了这里的5是x返回的时候就是X
}

func f3() (y int) {
	x := 5				//第二步 x=5
	defer func() {	
		x++				//第三步 修改的是x的值和y变量无关
	}()	
	return x			//第一步赋值  y=x=5
}
func f4() (x int) {
	defer func(x int) {
		// fmt.Println(&x)
		x++				//第二步 将返回的 x=5传入 x++这时候的x和外面的不再是同一个参数
		// fmt.Println(&x)
	}(x)
	// fmt.Println(&x)
	return 5			//第一步  赋值x=5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}