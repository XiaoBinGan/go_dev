package main
/**
 闭包的实际意义就是将传入的变量私有化暂存在栈中下次使用的时候还是继续使用的上次的结果值
*/
import (
	"fmt"
	"strings"
)


func add() func (int) int{
	var x int
	fmt.Printf("var X=%#d\n",x)
	return func (b int )int  {
	fmt.Printf("x+b=%#d\n",x+b)
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
	fmt.Printf("a(300):%#d \n",a(300))//+300
	fmt.Printf("a(500):%#d \n",a(500))//+500
	fmt.Printf("a(600):%#d \n",a(600))//+600
	fmt.Printf("a(600):%#d \n",a(600))//+600




	b :=suffix(".jpg")
	fmt.Println(b("sssss"))
	fmt.Println(b("ccccc"))
}