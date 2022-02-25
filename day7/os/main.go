package main

import (
	"fmt"
)



type student struct{
	Name  string 
	Age   int
	Score float32 
}

func main() {
	var str = "stu01 18 89.92"
	var stu student	
	fmt.Sscanf(str,"%s %d %f",&stu.Name,&stu.Age,&stu.Score)   
	fmt.Println(stu)
}
/*
%s string 输出
%d number 输出
%f float  输出
*/