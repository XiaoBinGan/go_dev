package main

import (
	"fmt"
	"runtime"
)

func main() {
	num :=runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	fmt.Println(num)
}