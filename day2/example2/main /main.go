package main
import (
	"fmt"
	"time"
)
const (
	man =1
	female=2
)

func main() {
	Second :=time.Now().Unix()////获取当前时间戳 单位是秒
	for{
		if(Second % female == 0){
			fmt.Println("female")
			}else if(Second % man == 0){
				fmt.Println("man")
			}
	}
	time.Sleep(100*time.Millisecond)

}