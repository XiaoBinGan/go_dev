package main
import "fmt"
func main (){
	str :="hellow wordl,中国"
	for i,v := range str {
		fmt.Printf("i[%d] v[%c] len[%d] \n",i,v,len([]byte(string(v))))
	}
}
//range的使用类似于each 
/*
%d代表输出整数
%c代表输出单个字符
%s代表输出字符串
/n表示换行
len表示一个数值的长度
**/