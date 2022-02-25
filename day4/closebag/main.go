package main
/**
 闭包的实际意义就是将传入的变量私有化暂存在栈中下次使用的时候还是继续使用的上次的结果值
*/
import (
	"strings"
	"fmt"
)


func add() func (int) int{
	var x int 
	return func (b int )int  {
		x+=b
		return x
	}
}

/**
判断传入的字符串是否是包含 suffixname的后缀，
不包含则加上  包含则直接返回
*/
func suffix(suffixname string) func (string) string  {
	return func (name string )string  {
		if strings.HasSuffix(name,suffixname) {
			return name
		}
		return name+suffixname
	}	

	
}





func main() {
	a := add()
	fmt.Println(a(300))
	fmt.Println(a(500))
	fmt.Println(a(600))
	fmt.Println(a(600))




	b :=suffix(".jpg")
	fmt.Println(b("sssss"))
	fmt.Println(b("ccccc"))
}