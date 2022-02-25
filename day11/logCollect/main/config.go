package main

import (
	"go_dev/day11/logCollect/tailf"
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
)

type Config struct{//初始化配置的参数
	logLevel string								//日志级别
	logPath string								//日志路径
	chanSize int								//管道大小
	kafkaAddr string							//kafka的地址或者集群
	collectconf []tailf.CollectConf				//tailf的设计自 每个元素都是一个追踪目标
	etcdAddr string								//etc的地址
	etcdKey  string
	}   
// 初始化一个全局的config
var (
		appConfig *Config	
) 
//迁移至tailf包
// type CollectConf struct{//初始化日志收集参数 
// 	LogPath string
// 	Topic string
// }

func loadCollectConf(conf config.Configer)(err error)  {//接收配置文件操作的对象的 用来获取相对应的配置文件

		var cc tailf.CollectConf
		cc.LogPath=conf.String("collect::log_path")
		if len(cc.LogPath)==0{
			err = errors.New("invalid collect::log_path")
			return
		}
		cc.Topic=conf.String("collect::topic")
		if len(cc.Topic)==0{
			err = errors.New("invalid collect::topic")
			return
		}
		appConfig.collectconf=append(appConfig.collectconf,cc)//将配collectconf对象放到总配置文件当中
		
		return
}

func loadConf(conftype,filename string)(err error)   {
	conf,err :=config.NewConfig(conftype,filename) //初始化可以读取配置文件啊的对象
	if err !=nil{
		fmt.Println("new config failed ,err:",err)
		return
	}

	appConfig=&Config{} //初始化一个配置对象准备给logs对象初始化使用[用来存储所以的配置对象-]
	appConfig.logLevel=conf.String("server::log_level")//获取配置文件中的日志级别
	if len(appConfig.logLevel) <1{//如果没有指定日志级别加一个默认值
		appConfig.logLevel="debug" 
	}
	appConfig.logPath=conf.String("server::log_path")//获取配置文件的日志存放路径
	if len(appConfig.logPath)<1{
		appConfig.logPath="../logs/access.log "
	} 
	appConfig.chanSize,err =conf.Int("collect::chanSize")//获取配置文件的日志存放路径
	if err!=nil{
		appConfig.chanSize = 100  //如果管道存放的位数配置文件为空的话就默认为100
		return
	} 
	appConfig.kafkaAddr=conf.String("kafka::server_addr")//获取kafka的配置文件
	if len(appConfig.kafkaAddr)==0 {
		// appConfig.kafkaAddr="127.0.0.1:9902"
		err = fmt.Errorf("invalid etcd addr")
		return
	}
	appConfig.etcdAddr=conf.String("etcd::addr")//获取etcd的地址
	if len(appConfig.etcdAddr)==0{
		err = fmt.Errorf("invalid etcd addr")
		return
	}
	appConfig.etcdKey=conf.String("etcd::configKey")//获取etcd的地址
	if len(appConfig.etcdKey)==0{
		err = fmt.Errorf("invalid etcd configKey")
		return
	}

	err  = loadCollectConf(conf)//初始化日志收集函数
	if err !=nil{
		fmt.Printf("load collect conf failed,err:%v\n", err)
		return
	}

	return
}