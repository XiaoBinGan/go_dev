package mylogger
//日志文件的通用配置
import (
	"runtime"
	"path"
	"fmt"
	"errors"
	"strings"
)

//LoggerLevel 定义日志级别的类型
type LoggerLevel uint16

const (
	//UNKNOW 0
	UNKNOW LoggerLevel = iota
	//DEBUG    1
	DEBUG
	//TRACE    2
	TRACE
	//INFO     3
	INFO
	//WARNING  4
	WARNING
	//ERROR    5
	ERROR
	//FATAL    6
	FATAL
)

//为了防止用户输入的值不合法所以把用户输入的级别进行强转
func parseLogLevel(s string) (LoggerLevel, error) {
	s = strings.ToUpper(s)
	switch s {
	case "DEBUG":
		return DEBUG, nil
	case "TRACE":
		return TRACE, nil
	case "INFO":
		return INFO, nil
	case "WARNING":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		return UNKNOW, errors.New("无效的日志级别")
	}
}
//根据等级返回相对应的字符串
func getLogString(level LoggerLevel ) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOW"
	}
}

//获取执行日志的行数和名称以及执行函数的文件名称
//传入参数告诉函数是哪一层函数执行的
func getinfo(skip int) (funcName, fileName string, lineNo int) {
	pc,file,lineNo,ok := runtime.Caller(skip)//runtime是一个标准库拿到是谁调用了 文件名称 函数名称 执行行数 是否执行成功
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name() //获取执行了该函数那一层的名称
	funcName =strings.Split(funcName,".")[1]//取到的名称通常如"main.main" 所以拆分截取后面一位
	fileName = path.Base(file)              //取出文件名称剔除全路径
	return
}
