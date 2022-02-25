package main

import (
	"io"
	"net/http"
	"fmt"
	"text/template"
)

type Person struct {
	Name  string
	Title string
	Age   int
}

type Result struct {
	output string
}
//Result实现Write接口
func (r *Result)Write(b []byte) (n int,err error) {
	fmt.Println("called by template")
	r.output +=string(b)
	return len(b),nil
}

//申明一个myTemlater 的指针指向供全局使用
var myTemplate *template.Template
// 封装初始化模版动作
func initTemlate(filename string)(err error)  {
	myTemplate,err =template.ParseFiles(filename)
	if err !=nil{
		fmt.Println("parse file err:",err)
		return
	}
	return
}
func userInfo(w http.ResponseWriter,requset *http.Request) {
	fmt.Println("handle hello")
	var arr []Person //创建一个person数组
	p := Person{Name: "Mary001", Age: 10, Title: "我的个人网站"}
	p1 := Person{Name: "Mary002", Age: 10, Title: "我的个人网站"}
	p2 := Person{Name: "Mary003", Age: 10, Title: "我的个人网站"}
	arr = append(arr,p)
	arr = append(arr,p1)
	arr = append(arr,p2)
	
	resultWriter :=&Result{} //申明一个Result的内存地址
    //Result实现了writer接口所以写入规则更具自定义的方式写入到Result.oupt当中去
	io.WriteString(resultWriter,"hellow world")
	err :=myTemplate.Execute(w,arr)                                                                                                                                                                                                                                                                                                                                                                                                         
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("template render data:",resultWriter.output)




}
func main() {
	initTemlate("./index.html")
	http.HandleFunc("/user/info",userInfo)
	err := http.ListenAndServe("0.0.0.0:8080",nil)
	if err!=nil {
		fmt.Println("http listen failed")
	}
}






















// Execute applies a parsed template to the specified data object,
// writing the output to wr.
// If an error occurs executing the template or writing its output,
// execution stops, but partial results may already have been written to
// the output writer.
// A template may be executed safely in parallel, although if parallel
// executions share a Writer the output may be interleaved.
// Execute将解析后的模板应用于指定的数据对象，
//将输出写入wr。
//如果在执行模板或写入输出时发生错误，
//执行停止，但部分结果可能已经被写入
//输出写入器。
//模板可以安全地并行执行，即使是并行的
//执行共享一个写入器，输出可以交错。
	// func (t *Template) Execute(wr io.Writer, data interface{}) error {
	// 	if err := t.escape(); err != nil {
	// 		return err
	// 	}
	// 	return t.text.Execute(wr, data)
	// }
