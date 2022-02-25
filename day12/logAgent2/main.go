package main

import (
	"go_dev/day12/logAgent2/conf"
	"go_dev/day12/logAgent2/etcd"
	"go_dev/day12/logAgent2/kafka"
	"go_dev/day12/logAgent2/taillog"
	"go_dev/day12/logAgent2/utils"
	"sync"

	"fmt"
	"time"
)
var (
	cfg =new (conf.AppConf)
)
// func run()  {
// 	for{
// 		select {
// 		case line := <-taillog.Readchan():
// 			kafka.SendToKafka(cfg.KafkaConf.Topic,line.Text)
// 		default:
// 			time.Sleep(time.Second)
// 		}
// 	}
// }
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
	
func main() {
	/*1.读取ini文件的配置项*/
	err := ini.MapTo(cfg,"./conf/conf.ini")
	if err!=nil{
		fmt.Printf("read ini config failed ,err:%v\n",err)
		return
	}
	// fmt.Printf("%#v\n",cfg)
	fmt.Printf("read ini config success ...\n")
	/*2.init kafka connection*/
	err = kafka.Init([]string{cfg.KafkaConf.Address},cfg.KafkaConf.ChanMaxSize)
	if err!=nil{
		fmt.Printf("connect kafka failed,err:%v",err)
		return
	}
	fmt.Printf("kafka init success ...\n")
	/*2.1 get out bound ip*/
	ipstr,err :=utils.GetOutboundIp()
    if err!=nil{
		panic(err)
	}
	etcdConfKey :=fmt.Sprintf(cfg.EtcdConf.Key,ipstr)
	/*3.init etcd client
		3.1. get etc config value
	*/
	err =etcd.Init([]string{cfg.EtcdConf.Address},time.Duration(cfg.EtcdConf.TimeOut)*time.Second)
	if err!=nil{
		fmt.Printf("etcd init client failed ,err:%v\n", err)
		return
	}
	fmt.Printf("etcd connect success ...\n")
	//use ini.comf variable .EtcdConf.Key get value
	LogEntryConf,err :=etcd.GetConf(etcdConfKey)
	if err!=nil{
		fmt.Printf("get key failed err:%v\n",err )
		return
	}
	for i,v  := range LogEntryConf {
		fmt.Printf("index:%v, value:%+v\n",i,v)
	}
	/*4.collect log send to kafka
	4.1 range every logcollecter to create 
	4.1 send massage to kafka
	*/
	taillog.Init(LogEntryConf)
	
	/*taillog out of NewConfChan*/
	newconf :=taillog.NewConfChan()
	/*create sync waiteGroup waite the goroutine end*/
	var wg sync.WaitGroup 
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey,newconf)
   	wg.Wait()




	// //2.init tail to watch log file
	// err =taillog.InitTail(cfg.TaillogConf.FileName)
	// if err!=nil{
	// 	fmt.Printf("init tail failed ,err:%v\n",err)
	// 	return
	// }
	// fmt.Printf("tail file  init success ...\n")
	// run()	
}