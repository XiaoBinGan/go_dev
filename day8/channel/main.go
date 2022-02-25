package main

import (
	"fmt"
)
type St struct {
	name string
}
func main() {
	var stChan chan interface{}           //创建一个channel 类型为
	stChan = make(chan interface{},2)     //make(chan 接口类型,长度2)超过就报错
	t :=St{name:"Alben"}				  //创建一个类型St的变量
	stChan <- &t						  //将地址通过管道存入stchan中
	fmt.Println(stChan)

	var a interface{}					  //申明一个类型为interface类型的变量 
	a = <- stChan						  //a为interface类型的所以能接受任何类型的值
	
	var b *St							  //申明一个类型为指针St的类型的变量来接收
	b,ok :=a.(*St)						  //将a强转成指针St的类型
	if !ok {							  //如果失败的话则不ok 显示失败的原因
		fmt.Println("can not convert")
	}
	fmt.Println(b)
}