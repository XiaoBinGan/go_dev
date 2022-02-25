package main

import (
	"fmt"
)


type animal struct{
	name string
}
func (a animal) move ()  {
	fmt.Printf("%s在动",a.name)
}
type dog struct{
	feet uint8
	animal //这样dog就继承了animal的方法
}
func (d dog)wang()  {
	fmt.Printf("%s在汪汪汪的叫",d.name)
}

func main() {
	d1 :=dog{
		feet:4,
		animal:animal{name:"旺财"},
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()
}