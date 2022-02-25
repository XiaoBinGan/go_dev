package main

import (
	"fmt"
	"go_dev/logAgent/conf"
	"go_dev/logAgent/kafka"
	"go_dev/logAgent/taillog"
	"time"
)
var (
	cfg =new (conf.AppConf)
)
func run()  {
	for{
		select {
		case line := <-taillog.Readchan():
			kafka.SendToKafka(cfg.KafkaConf.Topic,line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	//0.read config file (ini)
	// cfg,err :=ini.Load("./conf/conf.ini")
	// if err!=nil{
	// 	fmt.Printf("load config file failed ,err:%v\n",err)
	// 	return
	// }		
	// 						// [kafka]    读取子类型
	// fmt.Println(cfg.Section("kafka").Key("address"))//127.0.0.1
	// fmt.Println(cfg.Section("kafka").Key("topic"))	//my_log
	// fmt.Println(cfg.Section("taillog").Key("path"))	//./my.log
	//1.init kafka connection
	err := ini.MapTo(cfg,"./conf/conf.ini")
	if err!=nil{
		fmt.Printf("read ini config failed ,err:%v\n",err)
		return
	}
	// fmt.Printf("%#v\n",cfg)
	fmt.Printf("read ini config success ...\n")
	err = kafka.Init([]string{cfg.KafkaConf.Arrdress})
	if err!=nil{
		fmt.Printf("connect kafka failed,err:%v",err)
		return
	}
	fmt.Printf("kafka init success ...\n")
	//2.init tail to watch log file
	err =taillog.InitTail(cfg.TaillogConf.FileName)
	if err!=nil{
		fmt.Printf("init tail failed ,err:%v\n",err)
		return
	}
	fmt.Printf("tail file  init success ...\n")
	run()	
}