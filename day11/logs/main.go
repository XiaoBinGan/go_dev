package main

import (
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego/logs"
)


func main() {
	config :=make(map[string]interface{})
	config["filename"]="./logcollect.log"   //日志文件地址
	config["level"]=logs.LevelDebug			//日志文件打印等级


	configStr,err :=json.Marshal(config)	//初始化成一个json数组
	if err!=nil{
		fmt.Println("marshal failed ,err:",err)
		return
	}
	logs.SetLogger(logs.AdapterFile,string(configStr)) //接受一个file 和初始化的字符串
	logs.Debug("this is a test, My name is %s","stu01")
	logs.Warn("this is a trace,my name is %s","stu02")
	logs.Trace("this is a warn, my name is %s","Alben")
}