package main

import (
	"io/ioutil"
	"bufio"
	"io"
	"fmt"
	"os"
)

// 使用for循环读取文件中的所有数据。
func readFromFile0(){
	fileobj,err :=os.Open("./main.go")
	for{	
		if err!=nil{
			fmt.Printf("open file failed,err:%v",err)
			return
		}
		//记得关闭文件
		defer fileobj.Close()
		//读取文件
		// var tmp = make([]byte, 188)   //指定读取的长度
		var tmp [188]byte
		n,err :=fileobj.Read(tmp[:])//截取数组切片
		if err == io.EOF{
			fmt.Println("读完了")
			return
		}
		if err!=nil{
			fmt.Println("文件读取失败")
		}
		fmt.Printf("读取了%d个字节",n)
		fmt.Println(string(tmp[:n]))//拿到字符串切片转换成字符串
		if n<188 {
			return
		}
	}
}
// bufio按行读取示例
// bufio是在file的基础上封装了一层API，支持更多的功能。
func readFromFile1()  {
	file,err :=os.Open("./main.go")
	if err !=nil{
		fmt.Println("打开文件失败,err:",err)
		return
	}
	defer file.Close()
	reader :=bufio.NewReader(file)
	for{
		line,err :=reader.ReadString('\n') //按行读取
		if err == io.EOF{//这里只是单行读完了
			if len(line) !=0{//每行的字符长度不小于1
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err!=nil{
			fmt.Println("文件读取失败")
			return
		}
		fmt.Print(line)
	}
}
// io/ioutil包的ReadFile方法能够读取完整的文件，只需要将文件名作为参数传入。
func readFromFile2()  {
	file,err :=ioutil.ReadFile("./main.go")
	if err!=nil{
		fmt.Println("read file failed,err:",err)
		return
	}
	//基本上都是将字符数组的切片转成字符串
	fmt.Println(string(file))
}

func main() {
	readFromFile0()
	readFromFile1()
	readFromFile2()
}