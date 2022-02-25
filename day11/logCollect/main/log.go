package main

import (
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego/logs"
)


//将配置文件的等级映射成logs对象的当中
func convertLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	default:
		return logs.LevelDebug
	}
}




//初始化日志配置
func initLogger() (err error) {
	//初始化一个map
	config :=make(map[string]interface{})
	config["filename"]=appConfig.logPath
	config["level"]=convertLogLevel(appConfig.logLevel)
	configStr,err :=json.Marshal(config)//将配置转换成json数组
	if err!=nil{
		fmt.Println("inintlogger faild ,mashal failed ,err:",err)
		return
	}
	logs.SetLogger(logs.AdapterFile,string(configStr))//Setlogger(string,sring)
	return
}