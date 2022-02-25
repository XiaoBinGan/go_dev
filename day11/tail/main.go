package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail" //追踪并读取文件
)

func main() {
	filename := "./my.log" //引入文件
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen: true, //重新打开文件
		Follow: true,//是否跟随
		Location:&tail.SeekInfo{Offset:0,Whence:2},//文件从什么位置开始读取
		MustExist: false,//文件不存在不报错
		Poll:      true,
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}
	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Println("tail fail close reopen,filenam:%\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg", msg)
	}

}
