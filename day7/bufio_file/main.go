package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	file,err :=os.Open("./test.log")
	if err !=nil {
		fmt.Println("read file err :",err)
		return
	}
	defer file.Close()
	reader :=bufio.NewReader(file)
		str, err :=reader.ReadString('\n')//只能读取一行 结束就非
	if err !=nil {
		fmt.Println("read string failed, err:",err)
		return 
	}
	fmt.Printf("read string,ret:%s\n", str)
}