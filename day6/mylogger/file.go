package mylogger

import (
	"time"
	"fmt"
	"path"
	"os"
)


//在文件里面与日志代码相关

//FileLogger struct
type FileLogger struct{
	Level     LoggerLevel		//日志级别
	filepath  string			//日志路径
	fileName  string			//文件名称  error.log access.log
	fileObj   *os.File			//打开一个文件记录日志文件  拿到结构体指针   os是包file是结构体所以这里的类型是*os.File
	errObj	  *os.File			//打开一个文件记录错误日志	
	maxFileSize  int64			//文件最大切割阈值
}

//NewFileLogger constructe 
func NewFileLogger(levelStr,fp,fn string,maxSize int64)*FileLogger{
	Loglevel, err :=parseLogLevel(levelStr)//将传入的字符串转换成相对应的日志等级
	if err!=nil{
		panic(err)
	}
	fl :=&FileLogger{//返回日志对象地址
		Level:Loglevel,
		filepath:fp,
		fileName:fn,
		maxFileSize:maxSize,
	}
	fileErr :=fl.initFile()
	if fileErr!=nil{
		panic(fileErr)
	}
	return fl
}
//拿结构体的日志等级来对应日志级别
func (l *FileLogger) enable(logLevel LoggerLevel) bool {
	// fmt.Println(logLevel <=l.Level)
	return logLevel <= l.Level //说明日志级别包含当前的级别
}
//文件操作初始化文件根据传入的地址和名称打开日志文件|或创建日志文件进行操作
func (l *FileLogger)initFile()error {
	path :=path.Join(l.filepath,l.fileName)//初始化文件路径
	fileObj,err :=os.OpenFile(path,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)//初始化文件对象*os.File
	if err!=nil{
		fmt.Println("open file failed err:",err)
		return err
	}	
	errObj,err :=os.OpenFile(path+".error",os.O_CREATE|os.O_CREATE|os.O_WRONLY,0644)//error以上级别的日志就直接用传入的日志名称加上error
	if err!=nil{
		fmt.Println("open file failed err:",err)
		return err
	}	
	l.fileObj=fileObj//将初始化好的对象给构造函数初始化日志对象
	l.errObj=errObj
	return nil
}
//关闭操作的文件不然报错的话可能会导致堆栈内存溢出
func (l *FileLogger)close()  {
	l.fileObj.Close()
	l.errObj.Close()
}
//文件按设定大小切割		需要传入*os.file对象指针
func (l FileLogger)checkSize(file *os.File)bool {
	FileInfo,err :=file.Stat()//拿到fileinfo接口
	if err!=nil{
		fmt.Println("get file size failed ,err:",err)
	}
	return l.maxFileSize>=FileInfo.Size()//用设定好的最大值对比获取到的文件值超了说明是需要分割的
}
//文件达到阀值的时候重命名做备份,关闭,然后打开新的文件,返回给结构体继续去写入
func (l FileLogger)sliceFile(file *os.File)(*os.File,error)  {
	
	// 1.备份文件 rename xx.log->xx.log.bak2020100912
	// 2.关闭原来的文件
	// 3.打开一个新的日志文件
	

	//获取原来的文件名称 ,获取文件的全路径,拼接时间戳bak后缀 然后打开新的文件继续操作
	now :=time.Now().Format("20060102150405")//获取当前的秒数
	fileInfo,err :=file.Stat()//获取fileInfo接口
	if err!=nil{
		fmt.Println("get fileinfo failed ,err:",err)
		return nil,err
	}
	oldName :=path.Join(l.filepath,fileInfo.Name())
	newName :=fmt.Sprintf("%s.bak%s\n",oldName,now)
	os.Rename(oldName,newName)
	file.Close()
	fileObj,err :=os.OpenFile(oldName,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err!=nil{
		fmt.Println("open file failed ,err:",err)
		return nil,err
	}
	return fileObj,nil
}
//将日志用户初始化的构造函数的等级传入, 再将格式化内容和参数传入  参数可有可无
func (l *FileLogger)log(lv LoggerLevel,format string,s ...interface{}) {
	if l.enable(lv) {
		msg :=fmt.Sprintf(format,s...)//format接收是格式 ,  s接收的是多个或者一个值然后在格式化到format当中导出mgs
		level :=getLogString(lv)
		now :=time.Now()//获取当前时间
		funcName,fileName,lineNo :=getinfo(3)//获取执行文件的名称和函数名称,还有执行行数
		if l.checkSize(l.fileObj){//大于等于阀值就需要切割
			newFile,err :=l.sliceFile(l.fileObj)//调用函数传入当前操作的文件重命名,并且返回一个新的文件进行日志存储
			if err!=nil{
				fmt.Println("create new file failed ,err:",err)
				return
			}
			l.fileObj=newFile
		}
		fmt.Fprintf(l.fileObj,"[%s][%s][%s %s %d] %s\n",now.Format("2006-01-02 15:04:05"),level,fileName,funcName,lineNo,msg)
		if lv>=ERROR{
			if l.checkSize(l.errObj){
				newErrFile,err :=l.sliceFile(l.errObj)
				if err!=nil{
					fmt.Println("creact newErrFile failed ,err:",err)
				}
				l.errObj=newErrFile
			}
			fmt.Fprintf(l.errObj,"[%s][%s][%s %s %d] %s\n",now.Format("2006-01-02 15:04:05"),level,fileName,funcName,lineNo,msg)
		}
	}
}
//Debug level-log		 1用户希望的格式	2用户传入的参数
func (l *FileLogger) Debug(format string,s ...interface{}) {
		l.log(DEBUG,format,s...)
}
//Trace level-log
func (l *FileLogger) Trace(format string,s ...interface{}) {
		l.log(TRACE,format,s...)
}
//Info level-log
func (l *FileLogger) Info(format string,s ...interface{}) {
		l.log(INFO,format,s...)
}
//Warning level-log
func (l *FileLogger) Warning(format string,s ...interface{}) {
		l.log(WARNING,format,s...)
}
//Error level-log
func (l *FileLogger) Error(format string,s ...interface{}) {
		l.log(ERROR,format,s...)
}
//Fatal level-log
func (l *FileLogger) Fatal(format string,s ...interface{}) {
		l.log(FATAL,format,s...)
}