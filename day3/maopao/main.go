package main

import "fmt"

func print (n int){
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("A")	
		}
		fmt.Println()
	}


}

func main(){
	print(7)	
}