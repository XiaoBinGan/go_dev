package main

import (
	"go_dev/day11/logCollect/kafka"
	// "time"
	"go_dev/day11/logCollect/tailf"
	"github.com/astaxie/beego/logs"
	"fmt"
)




func main() {
	// filename :="../conf/logagent.conf" //配置文件地址
	filename :="/Users/wujiahao/study&demo/GO/package/src/go_dev/day11/logCollect/conf/logagent.conf" //配置文件地址
	//配置文件初始化
	err :=loadConf("ini",filename)//初始化配置文件	
	if err!=nil {
		fmt.Printf("load conf failed,err%v\n", err)
		panic("load conf failed")
	}
	logs.Debug("config intiallize success")
	//日志收集初始化
	err  = initLogger()//日志收集初始化
	if err!=nil{
		fmt.Printf("load logger failed,err:%v\n", err)
		panic("load logger failed")
	}
	logs.Debug("log intiallize success")


	//etcd的初始化
	collectConf,err :=initEtcd(appConfig.etcdAddr,appConfig.etcdKey)
	if err!=nil{
		logs.Error("init etcd failed,err:%v",err)
		return
	}
	logs.Debug("etcd intiallize success")

	//tailf初始化
	// err =tailf.InitTail(appConfig.collectconf,appConfig.chanSize)//追踪所有的日志文件集群,和存放消息的管道大小配置传入操作
	err =tailf.InitTail(collectConf,appConfig.chanSize)//追踪所有的日志文件集群,将etcd获取的参数传入tailf进行追踪
	if err !=nil {
		logs.Error("init tail failed ,err:%v",err)
		return
	}
	logs.Debug("tailf intiallize success")

	//kafka初始化
	err =kafka.InitKafka(appConfig.kafkaAddr)
	if err!=nil{
		logs.Error("init kafka failed,err:%v",err)
		return
	}
	logs.Debug("kafka intiallize success")




	// go func ()  {//测试日志文件是不是能追踪和写入
	// 	var count int
	// 	for {
	// 		count++
	// 		logs.Debug("test for logger %d",count)
	// 		time.Sleep(time.Millisecond*1000)
	// 	}	
	// }()
	err = serverRun()
	if err!=nil{
		logs.Error("serverRun failed ,err:%v",err)
		return
	}

	logs.Debug("intiallize success")
	logs.Debug("load conf succ,config:%v",appConfig)
}