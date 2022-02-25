package main

import(
	"time"
	// "fmt"
	mylogger "go_dev/day6/mylogger"
)

//测试我们自己写的日志库
func main() {
	// log :=mylogger.NewLog("DEBUG")
	log :=mylogger.NewFileLogger("error","./","testlog",10*1024)
	for{
		log.Debug("这是一条debug日志,id:%d,name:%s",100,"string")
		// log.Error("这是一条debug日志")
		log.Error("这是一条error日志")
		// log.Error("这是一条debug日志")
		// log.Error("这是一条debug日志")
		// log.Error("这是一条debug日志")
		// log.Error("这是一条debug日志")
		// log.Error("这是一条debug日志")
		// log.Error("这是一条debug日志")
		// log.Error("这是一条debug日志")
		// log.Error("这是一条debug日志")
		// log.Debug("这是一条debug日志")
		// log.Debug("这是一条debug日志")
		// log.Debug("这是一条debug日志")
		// log.Debug("这是一条debug日志")
		// log.Debug("这是一条debug日志")
		// log.Debug("这是一条debug日志")
	    time.Sleep(time.Second)
	}

	
}