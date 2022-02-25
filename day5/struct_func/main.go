package main

import (
	"fmt"
)


type person struct{
	name string
	age  int
}


//构造函数:约定是new开头
//返回的结构体还是结构体指针
//当结构体比较大的时候尽量使用结构体指针.减少程序的内存开销
func newPerson(name string , age int ) *person {
	return &person{
		name:name,
		age:age,
	}
	
}


func main() {
p1 :=newPerson("张三",19)
p2 :=newPerson("李三",20)
fmt.Println(p2,p1)
}