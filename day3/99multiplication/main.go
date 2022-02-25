package main

import "fmt"

func multiplication(){
	for i := 0; i < 9; i++ {
		for j := 0; j <=i ; j++ {
			fmt.Printf("%d*%d=%d\t",(j+1),(i+1),(i+1)*(j+1))
		}
		fmt.Println("")
	}
}




func main() {
	multiplication()
}