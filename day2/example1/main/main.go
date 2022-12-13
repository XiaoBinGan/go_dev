package main

import (
	"fmt"

	calc "../add"
)

func main() {
	sum := calc.Add(200, 300)
	fmt.Println(sum)
}
