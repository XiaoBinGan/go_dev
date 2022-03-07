package main

import "fmt"

func main()  {
 /**
    1.map复习点
       申明方式
  */
    //map的申明方式1
    var a map[string]string
    a=make(map[string]string,2)
    a["name"]="1"
    a["name1"]="2"
    a["name2"]="3"
    delete(a,"name2")
    fmt.Printf("a:%v \n",a)
    fmt.Printf("type of a:%T\n", a)
    //map声明方式2
    var b = map[string]string{
        "name":"张三",
    }
    b["name1"]="lisi"
    fmt.Printf("b:%v \n",b)
    fmt.Printf("type of b:%T\n", b)

}
