package main
import ("fmt")
func pipeGo() {
	pipe :=make(chan int,3)//make创建一个空间 chan管道内存储int类型的值，限制为三个
	pipe <- 1
	pipe <- 2
	pipe <- 3
	
	
	
	pipe <- 4

	fmt.Println(pipe)

}