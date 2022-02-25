package main
import(
	"fmt"
	"time"
	);

func add(a int ,b int) int {
	var sum int 
	sum=a+b
	fmt.Println(sum)
	return sum 
}


func main() {
		add(200,300)
		go testGoroutego(400,500)

		fmt.Println("start goroute")
		go pipeGo()
		fmt.Println("end goroute")
		time.Sleep(10*time.Second)



		fmt.Println("输出")
}
//go run 文件名