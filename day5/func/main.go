package main

import (
	"fmt"
)

type St struct{
	Name  string
	Age   int
	Score int64 
}
/*
		相当于java的set和get方法
		因为p是复制出来的所以
		下面的方法Get方法没办法直接获取到P
		这里将指针传入这样外部申明的参数才能被赋值并改变  get方法才能获取到被改变的P
	    下面的St如果不加*则是一个值类型P则是拷贝了份外部传入的变量然后进行的操作外部则不会改变
*/
func (p *St) int(name string, age int,score int64)  {
	p.Name=name
	p.Age=age
	p.Score=score	
}

func (p St) get() St {
	return p
}    

func main() {
	var a St
	a.int("c",18,100)  				// 方法1（go帮你简化了调用的，直接点就可直接调用）
	// (&a).int("c",18,100)  		// 方法2 上面的方法完整的调用应该是使用取址符拿到变量存储的地址再进行调用
	b :=a.get()
	fmt.Println(b)


}