package main

import (
	"fmt"
)

func main() {
	//1.数组申明和初始化
	//数组的长度必须是常量或者固定值,注意长度是数组类型的一部分.一旦定义不能变
	//[3]int 和[2]不是一种类型
	var a [3]int
	var b [2]int
	a = [3]int{1,2,3,}
	b = [2]int{4,5,}
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	//我们还可以使用指定索引值的方式来初始化数组
	a1 :=[...]int{2,2,33,3,4,3,21,1,}
	fmt.Printf("this is a1'value :%v\n",a1)


	//3.多维数组
	//其他语言数组间隔用逗号分隔 go语言是用空格分隔
	//[[1 2] [3 4] [5 6]]
	var a2 [3][2]int
	 a2 = [3][2]int{
		 [2]int{1,2},
		 [2]int{3,4},
		 [2]int{5,6},
	 }
	 fmt.Println(a2)



	 //4.数组的遍历
	 // 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	// 方法2：for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
	//5.遍历二维数组
	for _, v := range a2 {
		for _, v1 := range v {
			fmt.Printf("%d\t",v1)
		}
		fmt.Println()
	}
}