package main

import (
	"go_dev/day11/logCollect/kafka"
	"fmt"
	"time"
	"github.com/astaxie/beego/logs"
	"go_dev/day11/logCollect/tailf"
)

// msg *tailf.TextMsg


func serverRun() (err error) {
	for {//死循环不断的去读日志 
		msg := tailf.GetOneLine()//获取一条日志记录
		err =sendToKafka(msg)   //j
		if err!=nil{
			logs.Error("send to kafka failed ,err:%v",err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}
//将kafka需要的参数传入
// type TextMsg struct {//每条追踪到的日志记录
// 	Msg string   //实际的每条massage
// 	Topic string  //存放到那个topic
// } 
func sendToKafka(msg *tailf.TextMsg)(err error)  {
	fmt.Printf("read msg:%s ,topic:%s\n",msg.Msg,msg.Topic)
	//将每条去除的日志发送值kafka去消费
	err = kafka.SendToKafka(msg.Msg,msg.Topic)
	return
}