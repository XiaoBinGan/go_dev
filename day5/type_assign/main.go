package main

import (
	"fmt"
)




//类型断言0
func assign(a interface{}){
	fmt.Printf("%T\n",a)
	str,ok :=a.(string)//输入你希望的类型然后判断是否正确
	if !ok{
		fmt.Println("猜错了")
	}else{
		fmt.Println("传入的是一个字符串",str)
	}
}
// 类型断言1
func assign1(a interface{})  {
	switch t :=a.(type){//x.(type)需要配合switch使用
	case int:
		fmt.Printf("这是一个int类型值为%v\n",t)
	case bool:
		fmt.Printf("这是一个bool类型值为%v\n",t)		
	case string:
		fmt.Printf("这是一个string类型值为%v\n",t)
	default:
		fmt.Println("不支持的类型")
	}
}



func main() {
	assign1("string")	
}