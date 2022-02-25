package main

import (
	"io"
	"fmt"
	"os"
)

//CopyFile function
func CopyFile(dstName,srcName string)(written int64 ,err error)  {
	//以读取的方式打开资源文件
	srcfile,err :=os.OpenFile(dstName,os.O_CREATE|os.O_TRUNC|os.O_RDONLY,0666)
	if err!=nil{
		fmt.Println("open file failed,err :",err)
		return 
	}
	defer srcfile.Close()  //close操作最好是放在err后面执行如果file这个对象没有的时候就会报错
	dstfile,err :=os.OpenFile(dstName,os.O_CREATE|os.O_TRUNC|os.O_RDONLY,0666)
	if err!=nil{
		fmt.Println("open file is failed,err:",err)
		return
	}
	defer dstfile.Close()
	return io.Copy(dstfile,srcfile)//io.Copy()
}


func main() {
	_,err :=CopyFile("./ds.txt","./src.txt")
	if err!=nil{
		fmt.Println("copy file failed,err",err)
		return
	}
	fmt.Println("copy done!")
}