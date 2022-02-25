package main
import (
	calc "../add"
	"fmt"
)
func main() {
	sum := calc.Add(200,300)
	fmt.Println(sum)
}