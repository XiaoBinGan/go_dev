package main
import "fmt"

func control(){
	str :="hello world,中国"
	for i,v := range str {
		if i>2 {
			continue
		}
		if i>3 {
			break
		}
		fmt.Printf("index[%d] val[%c] len[%d]",i,v,len([]byte(string(v))))
	}



}

func main(){
	// ccf :="qweqweqwe"
	// fmt.Printf("%d%n",len([]byte(string(ccf))))
	control()
}