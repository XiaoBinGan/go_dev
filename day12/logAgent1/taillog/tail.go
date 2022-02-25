package taillog


import (
	"fmt"
	"github.com/hpcloud/tail" 
)
var (
	tailObj *tail.Tail
	/*LogChan return out use*/
	LogChan chan string
)

/*InitTail init tail
	param{
		filepath string
	}
	tail.Config{...}
	retrun err|nil
*/
func InitTail(filename string)(err error)  {
	config :=tail.Config{
		ReOpen: true, 									//重新打开文件
		Follow: true,									//是否跟随
		Location:&tail.SeekInfo{Offset:0,Whence:2},		//文件从什么位置开始读取
		MustExist: false,								//文件不存在不报错
		Poll:      true,
	}
	tailObj,err = tail.TailFile(filename,config)
	if err != nil {
		fmt.Println("tail file err:", err)
		return err
	}
	return nil
}

/*Readchan reade log from tail.Line*/
func Readchan()<-chan *tail.Line{
	 return tailObj.Lines
}
// func Readchan() *tail.Tail{
// 	 return tailObj
// }