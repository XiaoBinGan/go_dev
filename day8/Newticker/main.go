package main

import (
	"fmt"
	"runtime"
)

func main() {
	num :=runtime.NumCPU()
    fmt.Println(num)
}