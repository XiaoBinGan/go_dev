package calc

import (
	"fmt"
)




//包中的标识符:变量名\函数名\结构体\接口等 如果首字符是小写的 则表示私有的(只有当前的包里可以使用)
// 首字母大学的标识符才能被外部的包调用


func init()  {
	fmt.Println("我最先自动执行 比mian的init还早")
}
//Add is calcmethod
func Add(x,y int) int  {
	return x + y
}