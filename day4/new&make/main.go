package main

import (
	"fmt"
)






func test()  {
	s1 :=new([]int)
	// fmt.Println(s1)

	*s1 =make([]int,4)//初始化new出来的内存
	(*s1)[1]=123
	fmt.Println(s1)




	s2 :=make([]int, 10)
	fmt.Println(s2)

	s2[2] =200
	fmt.Println(s2)
}

func main() {
	test()
}