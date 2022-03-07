package main

import (
	"fmt"
)


//常量
//定义常量之后不能修改
const pi = 3.14
const (
	statusOK = 200
	notFound = 404
)
//批量声明常量时,如果某一行声明后没有赋值,默认是和上一行是一致的
const(
	n1 =100
	n2
	n3
)

//定义数量级
// 定义数量级 （这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。）
const (
	_ = iota 
	KB = 1 <<(10 * iota)
	MB = 1 <<(10 * iota)
	GB = 1 <<(10 * iota)
	TB = 1 <<(10 * iota)
	PB = 1 <<(10 * iota)

)
// iota是go语言的常量计数器，只能在常量的表达式中使用。

// iota在const关键字出现时将被重置为0。
// const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 
// 使用iota能简化定义，在定义枚举时很有用。
const (
	c1 = iota //0
	c2        //1
	c3        //2
	c4        //3
)
//使用_跳过某些值
const (
	v1 = iota //0
	v2        //1
	_
	v4        //3
)
// iota声明中间插队
const (
	b1 = iota   //0
	b2 = 100	//100
	b3 =iota	//2
	b4			//3
)

func main() {
		fmt.Println("n1:",n1)	
		fmt.Println("n2:",n2)	
		fmt.Println("n3:",n3)	
}