package main

import (
	"fmt"
)

func swap(a,b *int)(*int,*int)  {
	a,b= b,a
	fmt.Println("qweqwe",a,b)
	fmt.Println(*b)
	return a,b
}

func main()  {
	a,b :=3,4
	//fmt.Println(&a,&b)
	c,d :=swap(&a,&b)
	fmt.Println(c,d)
	fmt.Println(*c,*d)

	a = *c
	fmt.Println(d)
	fmt.Println(*d)
	b= *d
	fmt.Println(a,b)
}