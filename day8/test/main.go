package main

import (
	"fmt"
)
  
func main() {
	var t =&Student{
		Name:"Alben",
		Age:18,
	}
	fmt.Println(t)
	t.Save() 
	// fmt.Println(t.Name)
}
// go run student.go main.go 设计两个文件需要运行的时候  需要执行多个文件