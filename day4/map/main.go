package main

import (
	"fmt"
	"sort"
)
/**
1. map简介
key-value的数据结构，又叫字典或关联数组
 a. 声明
var map1 map[keytype]valuetype
         key值的类型 value的类型
var a map[string]string
var a map[string]int
var a map[int]string
var a map[string]map[string]string
```声明是不会分配内存的，初始化需要make```
*/
func testMap()  {//最好估算好容量避免程序在运行期间动态扩容
	 a :=make(map[string]string, 10)//map的申明方式1
	 a["qwe"]="123"
	 a["qwe"]="123"
	 a["qwe"]="143"
	 fmt.Printf("map的申明方式1:%#v\n",a)
	 var b =map[string]string{//map的申明方式2
		 "name":"张三",
	 }
	 b["user"]="admin"
	fmt.Printf("map的申明方式2:%#v\n",b)
     //申明一个值为map的变量 map使用必须初始化 不然会报Panic
	 c :=make(map[string]map[string]string,10)
	 //当值为一个map的时候一定要注意初始化之后再去使用
	 c["demo1"]=make(map[string]string)
	 c["demo1"]["demo_1"]="str1"
	 c["demo1"]["demo_2"]="str2"
	 c["demo1"]["demo_3"]="str3"
	 c["demo1"]["demo_4"]="str4"
	 c["demo1"]["demo_5"]="str5"
	 c["demo2"]=make(map[string]string)
	 c["demo2"]["demo_1"]="str1"
	 c["demo2"]["demo_2"]="str2"
	fmt.Printf("map的申明方式2的长度:%#d\n",len(c))
	fmt.Printf("map的申明方式2的值:%#v\n",c)




}
/**
5. map排序
a. 先获取所有key，把key进行排序
b. 按照排序好的key，进行遍历
*/
func mapsort()  {
	n :=map[string]int{
		"st1":123,
		"st2":312,
		"st3":534,
		"st4":745,
	}
	// for _, v := range n { golang是无序的
	// 	fmt.Println(v)
	// }
	var m []string
	for k:= range n {
		m=append(m,k)
	}
	sort.Strings(m) 
	for _,v := range m{
		fmt.Println(n[v],"_")
	}

	fmt.Println(n)
	fmt.Println(m)
}

//map反转初始化另外一个map，把key、value互换即可
func mapreserve()  {
	n :=map[string]int{
		"st1":123,
		"st2":312,
		"st3":534,
		"st4":745,
	}
	m :=make(map[int]string,4)
	for k, v := range n {
		m[v]=k
	}

	fmt.Println(n)
	fmt.Println(m)

}






func main() {
	testMap()
	//mapsort()
	//mapreserve()
}