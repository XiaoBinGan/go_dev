package main

import (
	"fmt"
	"strconv"
)
//只有一个strconv.p=Itoa是能将数值转换成字符串的  其他都是用来将字符串转换成 数值的
func main() {
	var a string ="100"
	Int,err :=strconv.Atoi(a)//Atoi()函数用于将字符串类型的整数转换为int类型
	if err!=nil{
		fmt.Println("can't convert to int")
	}else{
		fmt.Printf("type:%T,value:%d\n",Int,Int)
	}

	str1 :=strconv.Itoa(Int)//Itoa()函数用于将int类型数据转换为对应的字符串表示
	fmt.Printf("type:%T,value:%s\n", str1,str1)

	// Parse系列函数
	// Parse类函数用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()。
	var b string ="True"
	bol,err :=strconv.ParseBool(b)//接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；错误。
	if err!=nil{
		fmt.Println("can't parse string to bool,err:",err)
	}else{
		fmt.Println(bol)
	}

	PInt,err :=strconv.ParseInt("-2",10,64)// 返回字符串表示的整数值，接受正负号。
	if err!=nil{
		fmt.Println("con't parse int to string err:",err)
	}else{
		fmt.Printf("type:%T,value:%d\n",PInt,PInt)
	}

	PuIbt,err :=strconv.ParseUint("3",10,64)// ParseUint类似ParseInt但不接受正负号，用于无符号整型。
	if err!=nil{
		fmt.Println("con't parse int to string err:",err)
	}else{
		fmt.Printf("type:%T,value:%d\n",PuIbt,PuIbt)
	}

	r :="3.1415926"
	fl,err :=strconv.ParseFloat(r,64)// 解析一个表示浮点数的字符串并返回其值。
	if err!=nil{
		fmt.Println("can't parse string to float,err:",err)
	}else{
		fmt.Printf("type:%T,value:%f\n",fl,fl)
	}


}


// strconv包实现了基本数据类型与其字符串表示的转换，主要有以下常用函数： Atoi()、Itia()、parse系列、format系列、append系列。
// 这一组函数是我们平时编程中用的最多的。

//1. Atoi()
// Atoi()函数用于将字符串类型的整数转换为int类型，函数签名如下。

// func Atoi(s string) (i int, err error)


//2. Itoa()
// Itoa()函数用于将int类型数据转换为对应的字符串表示，具体的函数签名如下。
// func Itoa(i int) string

//3.ParseBool()
// func ParseBool(str string) (value bool, err error)
// 返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。

//4.ParseInt()
// func ParseInt(s string, base int, bitSize int) (i int64, err error)
// 返回字符串表示的整数值，接受正负号。

// base指定进制（2到36），如果base为0，则会从字符串前置判断，”0x”是16进制，”0”是8进制，否则是10进制；

// bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；

// 返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。

//5.ParseUnit()
// func ParseUint(s string, base int, bitSize int) (n uint64, err error)
// ParseUint类似ParseInt但不接受正负号，用于无符号整型。

// ParseFloat()
// func ParseFloat(s string, bitSize int) (f float64, err error)
// 解析一个表示浮点数的字符串并返回其值。

// 如果s合乎语法规则，函数会返回最为接近s表示值的一个浮点数（使用IEEE754规范舍入）。

// bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；

// 返回值err是*NumErr类型的，语法有误的，err.Error=ErrSyntax；结果超出表示范围的，返回值f为±Inf，err.Error= ErrRange。

