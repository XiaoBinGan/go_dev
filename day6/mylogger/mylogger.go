package mylogger

import (
	"fmt"
	"time"
)


//Logger log's struct
type Logger struct {
	Level LoggerLevel
}
//NewLog constructe
func NewLog(levelStr string) Logger {
	Level, err := parseLogLevel(levelStr) //根据用户输入的日志级别来构建Logger结构体
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: Level,
	}
}

//拿结构体的日志等级来对应日志级别
func (l Logger) enable(logLevel LoggerLevel) bool {
	// fmt.Println(logLevel <=l.Level)
	return logLevel <= l.Level //说明日志级别包含当前的级别
}

func (l Logger)log(lv LoggerLevel,format string,s ...interface{}) {
	if l.enable(lv) {	
		msg :=fmt.Sprintf(format,s...)//format接收是格式 ,  s接收的是多个或者一个值然后在格式化到format当中导出mgs
		level :=getLogString(lv)
		now :=time.Now()//获取当前时间
		funcName,fileName,lineNo :=getinfo(3)//获取执行文件的名称和函数名称,还有执行行数
		fmt.Printf("[%s][%s][%s %s %d] %s\n",now.Format("2006-01-02 15:04:05"),level,fileName,funcName,lineNo,msg)
	}	
}

//Debug level-log
func (l Logger) Debug(format string,s ...interface{}) {
		l.log(DEBUG,format,s...)
}
