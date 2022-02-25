package main
//go语言当中禁止循环导入包  a导入b b导入c c导入a
import (
	"fmt"
	calc "go_dev/day5/package1/calc" //这里的路径是基于gopath/src/的路径继续往下写的
)



//go语言的执行导入包的语句惠子自动触发包内部的init()函数的调用.
//需要注意的是:init()函数没有参数也没有返回值
//init()函数在程序运行时自动被调用执行,不能在代码中主动调用它
//但是全局的声明  var const 这种申明执行早于init函数
func init()  {
	fmt.Println("main的init函数执行")
}
func main() {
	ret := calc.Add(1, 2)
	fmt.Println(ret)
}
