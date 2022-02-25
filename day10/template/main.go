package main

import (
	"os"
	"fmt"
	"text/template"
)



// 1定一个struct
type Person struct{
	Name   string
	Title  string
	age    int
}

func main() {
	t,err :=template.ParseFiles("./index.html")
	if err!=nil {
		fmt.Println("pare file err:",err)
		return
	}
	p := Person{Name: "Mary", age:31, Title: "我的个人网站"}
	err =t.Execute(os.Stdout,p)//将p的值输出到模版 然后输出到控制台
	if err!=nil{
		fmt.Println("Tere was an error:",err.Error())
	}
}
// t,err :=template.ParseFiles("./index.html")
// 创建一个新模板并解析模板定义
// 指定的文件。返回的模板名称将包含基名称和
// 已解析第一个文件的内容。必须至少有一个文件。
// 如果出现错误，解析停止，返回的*模板为nil。
// //
// 解析不同目录中具有相同名称的多个文件时，
// 最后一个提到的将会是结果。
// 例如，ParseFiles("a/foo"， "b/foo")将"b/foo"存储为模板
// 命名为“foo”，而“a/foo”不可用。

// err :=t.Execute(os.Stdout,p)//将p的值输出到模版 然后输出到控制台
// Execute将解析后的模板应用于指定的数据对象，并将输出写入wr。
// 如果在执行模板或写入输出时发生错误，执行将停止，但可能已经将部分结果写入到输出写入器。
// 模板可以安全地并行执行，但如果并行执行共享一个写入器，则输出可以交错执行。
// 如果数据是反射。值时，模板应用于反映的具体值。Value保持不变，如fmt.Print。