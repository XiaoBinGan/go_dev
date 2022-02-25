package main

import (
	"fmt"
)
//St interface
type St interface {
	Run()
	start()
}
//Student struct
//一个对象只要全部实现了接口中的方法，那么就实现了这个接口。
//换句话说，接口就是一个需要实现的方法列表。
type Student struct {
	Name string
	Age  int
}
//Run methods
func (p *Student) Run() {
	fmt.Println(p)
}

func (p Student) start() {
	p.Name = "yhan"
	fmt.Println(p.Name)
}

func main() {
	var a St
	b := Student{
		Name: "alben",
		Age:  12,
	}
	// b.Run()
	// b.start()
	a = &b
	a.Run()
	a.start()
}
