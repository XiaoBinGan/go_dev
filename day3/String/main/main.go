package main

import (
	"strings"
	"fmt"
)
func urlProcess (url string) string {
	result :=strings.HasPrefix(url,"https://")//是否以什么字符串开头 返回值Boolean
	if !result {
		url =fmt.Sprintf("https://%s",url)
	}
	return url
}
func pathProcess (path string) string {
	result :=strings.HasSuffix(path,"/")//是否以什么字符串结尾 返回值Boolean
	if !result {
		path =fmt.Sprintf("%s/",path)
	}
	return path
}
func main() {
	var (
		url string
		path string
	)
	fmt.Scanf("%s%s",&url,&path)//未解释的字节的字符串或切片 &表示传入值得地址
	url=urlProcess(url)
	path=pathProcess(path)

	fmt.Println(url)
	fmt.Println(path)
	
}
 