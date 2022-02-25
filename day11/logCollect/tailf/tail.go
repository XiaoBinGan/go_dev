package tailf

import (
	"sync"
	"time"
	"github.com/astaxie/beego/logs"
	// "fmt"
	"github.com/hpcloud/tail"
)
const(
	StatusNormal = 1
	StatusDelete = 2
)

type CollectConf struct{//初始化日志收配置文件 
	LogPath string `json:"logpath"`
	Topic string   `json:"topic"`
}

type TailObj struct{//单个tail的配置文件
	tail *tail.Tail
	conf CollectConf 
	status int 			 //判断当前追踪的文件是否存在
	exitChan chan int    //退出的chan
}

type TextMsg struct {//每条追踪到的日志记录
	Msg string   //实际的每条massage
	Topic string  //存放到那个topic
}

type TailObjMgr struct{//tail实例集群
	tailobjs []*TailObj
	msgChan chan *TextMsg //将n条记录存放到chan当中然后kafka再去消费
	lock     sync.Mutex	  //可能有删除添加操作所以加一个互斥锁

}
var (
	tailobjMgr *TailObjMgr 
)
//初始化每个tail的配置
func InitTail(conf []CollectConf ,chanSize int)  (err error) {
	//初始化tail集群的收集器
	tailobjMgr =&TailObjMgr{
		msgChan :make(chan*TextMsg,chanSize),  
	}
	if len(conf)==0{//判断传入的参数是不是满足最小条数1
		logs.Error("invalid config for log collect,conf:%v",conf)
		return
	}
	// 遍历传入的配置文件的数组[]CollectConf生成tail对象
	for _, v := range conf {
		createNewTask(v)
	}
	return
}
//存放每个tail的记录
func ReadFromTail(tailobj *TailObj)  {//读取实时文件内的记录存放到chan当中
	for true {//循环条件
		line,ok :=<-tailobj.tail.Lines  //使用初始化好的tail对象读取每条记录
		if !ok {
			logs.Warn("tail file close reopne,filenam:%s\n",tailobj.tail.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}

		textmsg :=&TextMsg{//将每条记录存放到 TextMsg类型的struct当中
			Msg:line.Text,
			Topic:tailobj.conf.Topic,
		}
		tailobjMgr.msgChan<-textmsg//将每条记录塞到设定好的chan当中等待传入kafka

	}
}
//不需要参数  返回类型TextMsg
//输出每条
func GetOneLine()(msg *TextMsg)  {
	msg =<- tailobjMgr.msgChan  //读取一条chan里面的数据
	return
}
//实现etcd传入参数动态跟新tailf
func UpdateConfig(confs []CollectConf)(err error)  {
	tailobjMgr.lock.Lock()									//开始操作就锁死防止其他线程操作
	defer tailobjMgr.lock.Unlock()							//操作结束返回之后解锁
	for _, oneConf := range confs {							//遍历conf []CollectConf
		var isRunning bool = false						    //是否继续读取的标志位
		for _, obj := range tailobjMgr.tailobjs {			//遍历原先的配置文件然后对比etcd传入的配置
			if oneConf.LogPath == obj.conf.LogPath {	
					isRunning = true
					obj.status=StatusNormal					 //如果一样的话循环结束
					break
			}
		}
		if isRunning{										 //如果配置没有变化的就继续循环
			continue
		}
		createNewTask(oneConf)								 //初始化新的配置客户端跟踪日志文件 
	}
	var 	tailobjs []*TailObj
	for _, obj := range tailobjMgr.tailobjs {			     //遍历原先的配置文件然后对比etcd传入的配置
		obj.status = StatusDelete
		for _, oneConf := range confs {
			if oneConf.LogPath == obj.conf.LogPath {	
				obj.status=StatusNormal						 //如果一样的话循环结束
				break
			}
		}

		if obj.status == StatusDelete{ 
			obj.exitChan <- 1
			continue
		}
		tailobjs = append(tailobjs,obj)							//没被删除的内容添加到数组当中
	}
		tailobjMgr.tailobjs=tailobjs
	return
}
//发现etcd当中的配置有新增的配置项就调用creatNewTask创建新的tailf去追踪
//将[]CollectConf当中未被监听的一项
func createNewTask(conf CollectConf){
	obj :=&TailObj{//每循环一次生成一个TailObj对象接收
		conf:conf,
		exitChan: make(chan int,1),//初始化退出的chan
	}
   //单个配置的初始化,将每个v的filename 传入进行初始化然后生成tails对象然后去操作你想监控的日志文件		
	tails,errTail :=tail.TailFile(conf.LogPath,tail.Config{	 //初始化配置tail客户端
		ReOpen:true,
		Follow:true,
		MustExist:false,
		Poll:true,
	})
	if errTail!=nil{										//如果失败了就输出日志
		logs.Error("tailf config failed err:",errTail)
		return
	}
	obj.tail=tails											//将初始化的客户端的配置//将tail当前初始化传入TailObj
	tailobjMgr.tailobjs=append(tailobjMgr.tailobjs,obj)		//将初始化的一个tail放入集群中管理
	go ReadFromTail(obj)									
}