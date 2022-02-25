package taillog


import (
	"context"
	"go_dev/day12/logAgent2/kafka"
	"fmt"
	"github.com/hpcloud/tail" 
)
var (
	tailObj *tail.Tail
	/*LogChan return out use*/
	LogChan chan string
)
/*Tailtask a log collect task
  path  watch log path
  topic kafka save logs container
  instance a (tail.Tail) watch log client 
  close the tailTask so create the context object
  ctx,cancel := context.WithCancel(context.Background())
  cancelFunc:cancel
*/
type Tailtask struct{
	path string
	topic string
	instance *tail.Tail
	ctx context.Context
	cancelFunc context.CancelFunc

}
/*NewTailtask construct of a new T ailtask*/
func NewTailtask(path,topic string)(tailObj *Tailtask)  {
	ctx,cancel :=context.WithCancel(context.Background())
	tailObj=&Tailtask{
		path:path,
		topic:topic,
		ctx:ctx,
		cancelFunc:cancel,
	}
	tailObj.init()
	return
}
/*initTail init tail.Tail client for Tailtask
  send lins and to kafa's topic
*/
func (t *Tailtask)init()() {
	config :=tail.Config{
		ReOpen: true, 									//reOpen file
		Follow: true,									//following ?
		Location:&tail.SeekInfo{Offset:0,Whence:2},		//read from where
		MustExist: false,								//fail Does not exist without errors
		Poll:      true,
	}
	var err error
	t.instance,err = tail.TailFile(t.path,config)
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}
	go t.run()
}
/*run if context.Done not nil,stop the go routine 
select every line's massage send to kafka topic*/
func (t *Tailtask)run(){
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("推出的是%s_%s\n",t.path,t.topic)
			return
		case line :=<-t.instance.Lines:
			kafka.SendToChan(t.topic,line.Text)
		}
	}
}
