package main

import (
	"fmt"
)

func f1(x,y int) int {
	a :=x+y
	fmt.Println(a)
	return a
}

func f2(f func(int,int)int,m,n int) func() {
	tmp :=func ()  {
		f(m,n)
	}
	return tmp
}  


func main() {
	fmt :=f2(f1,100,20)
	fmt()
}