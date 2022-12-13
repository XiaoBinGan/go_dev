package main

import (
	"fmt"
)

func main() {
	    var b  int= 20
		if  b >18{
			fmt.Println("你成年了")
		}else{
			fmt.Println("你还没成年")
		}


		//作用域
		//age变量只在if调价判断语句的作用域当中生效出了这个作用域就拿到拿不到age了
		if age:=19;age>20 {
			fmt.Println("你成年了2")
		}else{
			fmt.Println(age)
		}
		// fmt.Println(age)这里会报错


}