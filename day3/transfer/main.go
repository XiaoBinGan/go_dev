package main
import "fmt"

func modify(a *int){ //  *a 存储单元的值  &a存储单元的地址
	*a = 100
}

func main() {
	a := 8
	// fmt.Println(a)
	fmt.Printf("%d \n",a)
	modify(&a)
	// fmt.Println(a)
	fmt.Printf("%#v \n",&a)
}
