package main

import (
	"fmt"
)

//链式调用的精髓就是返回当前的对象继续操作
type Stu struct{
	Name string
	Age  int 
}
func (p *Stu) SetName(name string) *Stu {
	p.Name=name
	return p
}

func (p *Stu) SetAget(age int) *Stu  {
	p.Age=age
	return p
}
func (p *Stu) Print() *Stu {
	fmt.Println("age:%d name:%s\n",p.Age,p.Name)
	return p
}

func main() {
	stu :=&Stu{}
	stu.SetAget(12).SetName("alben").Print()
}