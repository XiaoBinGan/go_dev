package main

import (
	"io/ioutil"
	"bufio"
	"fmt"
	"os"
)
/*
模式	含义
os.O_WRONLY	只写
os.O_CREATE	创建文件
os.O_RDONLY	只读
os.O_RDWR	读写
os.O_TRUNC	清空
os.O_APPEND	追加
perm：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。
*/

//WriteString  write demo
func WriteString()  {
	file,err :=os.OpenFile("./demo.txt",os.O_CREATE|os.O_TRUNC|os.O_RDWR,0666)
	if err!=nil{
		fmt.Println("open file failed,err:",err)
		return
	}
	defer file.Close()//这个步骤是操作文件不能少的关闭文件
	str :="Hello world! golang"
	file.Write([]byte(str))   //写入字节切片数据  将字符串转换成字节数组
	file.WriteString("hello 新世界")		  //直接传入字符串就可以了
}

//bufio.NewWriter
func bufioNewWiter() {
	file,err :=os.OpenFile("./demo.txt",os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0666)
	if err!=nil{
		fmt.Println("open file failed ,err:",err)
		return
	}
	defer file.Close()
	writer :=bufio.NewWriter(file)
	// for i := 0; i <10; i++ {
		writer.WriteString("Hello golang writer.writeString")//将数据先写入缓存
	// }
	writer.Flush() //将缓存中的内容写入文件
}
func ioWriterFild()  {
	str :="hello golang"
	err :=ioutil.WriteFile("./demo.txt",[]byte(str),0777)
	if err!=nil{
		fmt.Println("write file failed ,err:",err)
		return
	}

}

func main() {
	// WriteString()
	// bufioNewWiter()
	ioWriterFild()
}