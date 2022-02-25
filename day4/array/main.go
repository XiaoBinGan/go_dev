package main

import (
	"fmt"
)

/**
斐波那契额数列
*/
func feibo(n int){
   var a= make([]uint64,10)
	a[0]=1
	a[1]=1
	for i := 2; i < n; i++ {
		a[i]=a[i-1]+a[i-2]
	}
	for _, v := range a{
		fmt.Println(v)
	}
}


/**
 二维数字的初始化和展示
*/
func array() {
	//      x,y        x,y       x行数 y列数
	var arr[4][3]int=[...][3]int{
									{1,2,3},
									{4,5,6},
									{1,2,3},
									{4,5,6}}
	for _, v := range arr{

		for _, row := range v {
			fmt.Println(row,"_")
		}
	}
	// fmt.Println(arr)
}






func main() {
	feibo(10)
	array()
}