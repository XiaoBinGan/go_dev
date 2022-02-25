package main

import (
	"fmt"
)

func main() {
	// fmt.Print("不会换行")
	// fmt.Println("会换行")
	// // fmt.Printf("格式化字符串", 值)
	// 	        // %T :查看类型
	// 	        // %d :十进制数
	// 	        // %b :二进制数
	// 	        // %o :八进制数
	// 	        // %x :十六进制数
	// 			// %c :字符对应Unicode码值
	// 			// &U :表示为Unicode格式
	// 	        // %s :字符串
	// 	        // %p :指针
	// 	        // %f :浮点数
	// 			// %v :值
	// 			// %+v:输出结构体时会添加字段名
	// 			// %#v:详细的值
	// 			// %% :输出百分号
	// 			// %t :布尔值
	// var m1 = make(map[string]int, 1)
	// m1["理想"] =100
	// fmt.Printf("%v\n", m1)
	// fmt.Printf("%#v\n", m1)
	// fmt.Printf("%d%%\n",100)
	// fmt.Printf("%q\n", 65)//输出数值对应的字符表
	/*
		Go语言fmt包下有
		fmt.Scan、
		fmt.Scanf、
		fmt.Scanln
		三个函数，可以在程序运行过程中从标准输入获取用户的输入。
	*/
	// Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
	// 本函数返回成功扫描的数据个数和遇到的任何错误。如果读取的数据个数比提供的参数少，会返回一个错误报告原因。
	// var (
	// 	name string
	// 	age  int
	// 	married bool
	// )
	// fmt.Scan(&name, &age, &married)
	// fmt.Printf("用户输入结果 name:%s age:%d married:%t \n", name, age, married)
	// Scanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中。
	// 本函数返回成功扫描的数据个数和遇到的任何错误。
		var (
			name    string
			age     int
			married bool
		)
		fmt.Scanln("1:%s 2:%d 3:%t", &name, &age, &married)
		fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}