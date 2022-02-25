package main
import "fmt"
var a string 
func main() {
	var c int8 =100
	// var d int16=c这是一种错误的写法因为两者是不同的类型所以不能直接赋值需要做类型的转换
	var d int16 =int16(c)
	fmt.Printf("c=%d d=%d\n",c,d)

	a= "G"
	fmt.Println(a)
	f1()
}
func f1()  {
	a :="o"
	fmt.Println(a)
	f2()
}
func f2()  {//这里虽然是在函数内部调用但是执行的时候f2的内部并没有a变量所以指向的还是全局的a
	fmt.Println(a)
}
