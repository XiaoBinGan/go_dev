package main

import (
	"fmt"
)



type Cart1 struct{
	name string
	age  int
}
type Cart2 struct{
	name string
	age  int
}
type Cart3 struct{
	score int
}
type A struct{
	Cart1					//匿名嵌套结构体可直接调用被嵌套结构体内部的参数 但是要保证被嵌套的参数不能重复
	Cart2					//匿名嵌套
	Cart3					//匿名嵌套
	Page int 
	age int 
}
func main() {
	var t A
	t.Cart1.name="alben"
	t.Cart2.name="joyce"    
	t.Cart2.age=19          //结构体内部有相同参数名冲突的时候想调用下层匿名字段的参数赋值就得加上匿名参数的名称再去点调用
	t.score=100             //结构体内部如果不冲突可以直接匿名赋值给下一层的参数
	t.age=100				//结构体内部有这个值就直接赋值给最外层的结构体的参数
	fmt.Println(t) 
}
/*
匿名字段的调用
匿名字段的冲突处理
匿名字段和申明的对象内部的字段冲突了取结构内部的字段

*/
