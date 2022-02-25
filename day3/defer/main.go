 package main
import "fmt"
// 1. 当函数返回时，执行defer语句。因此，可以用来做资源清理
// 2. 多个defer语句，按先进后出的方式执行
// 3. er语句中的变量，在defer声明时就决定了。

func  a()  {
	i :=0 
	defer fmt.Printf("defer i %d\n",i)//3`
	defer fmt.Printf("defer i= %d\n",i)//2`
	i++
	fmt.Println(i)//1
	return 
}

func f()  {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("defer %d\n",i)//defer 后面的语句都会延迟执行后进先出
		//fmt.Printf("%d\n",i)
	}	
}



func main() {
	a()//1
	f()//4
}