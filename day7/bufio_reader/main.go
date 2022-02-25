package main

import (
	"bufio"
	"fmt"
	"os"
)
/*function description
		bufio.NewReader 创建缓冲区读取文件
		os
		Stdin、Stdout和Stderr是指向标准输入、标准输出、标准错误输出的文件描述符。
*/
func rd() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read string failed, err:", err)
		return
	}
	fmt.Printf("read str succ,ret:%s\n", str)
}







func main() {

}
