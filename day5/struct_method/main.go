package main

import (
	"fmt"
)
/*St 在go语言当中的首字母大写的标识符:变量名 函数名 类型名 方法名 
  就表示对外部包可见的(暴露的,公有的)
  所以go提示你必须要写相对的注释格式如下
  St空格this is St struct
*/
//St this is St struct
type St struct{
	Name  string
	Age   int
	Score int64 
}
//将结构体的内存地址传入这样就直接作用的是传入的这个结构体,而不是函数内部复制一份修改
func (s *St) int(name string, age int,score int64)  {
	s.Name=name
	s.Age=age
	s.Score=score	
}
//通常上面用了指针传入后面都应该用指针但是这里是为了学习所以拷贝一份传入的参数
func (s St) get() St {
	return s
}    

func main() {
	var a St
	a.int("c",18,100)  				// 方法1（go帮你简化了调用的，直接点就可直接调用）
	// (&a).int("c",18,100)  		// 方法2 上面的方法完整的调用应该是使用取址符拿到变量存储的地址再进行调用
	b :=a.get()
	fmt.Println(b)


}
/*
	Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。
	接收者的概念就类似于其他语言中的this或者 self。
	方法的定义格式如下：
		func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
			函数体
		}
	其中，

	接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，
	而不是self、this之类的命名。
	例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
	接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
	方法名、参数列表、返回参数：具体格式与函数定义相同。

*/