package main

import (
	"os"
	"io"
	"fmt"
	"bufio"
)


// 使用fallthrough强制执行后面的case代码

type count struct{
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}
func main() {
	 file , err :=os.Open("./test.log")
	 if err!=nil {
		 fmt.Println("read file err:",err)
	 }
	 defer file.Close()//切记一定要在结束之后关闭文件 
	 var  c count
	 reader :=bufio.NewReader(file)
	 for {
		str,err :=reader.ReadString('\n')//如果最后一行有内容会读不全，所以需要在结束的时候换多一行才能全部读取
		if err ==io.EOF{
			break
		   // fmt.Printf("read string end:%v",err)
		}
		if err !=nil{
			fmt.Printf("read string error:%s",err)
			return
		}
		fmt.Printf("this is a string:%s\n",str)
		runeArr :=[]rune(str)//强转成字符数字
		// runeArr1 :=[]byte(str)//强转成字节数字 
		// fmt.Println(runeArr)
		// fmt.Println(runeArr1)
		for _, v := range runeArr {
			switch {
			case v>='a'&& v<='z':
				fallthrough
			case v>='A'&& v<='Z':
				 c.ChCount++
			case v>=' '&& v<='\t':
				c.SpaceCount++
			case v>=0&& v<=9:
				c.NumCount++
			default:
				c.OtherCount++
			}
		}
	}
		fmt.Printf("char count:%d\n",c.ChCount)
		fmt.Printf("Number count:%d\n",c.NumCount)
		fmt.Printf("Space count:%d\n",c.SpaceCount)
		fmt.Printf("other count:%d\n",c.OtherCount)
}