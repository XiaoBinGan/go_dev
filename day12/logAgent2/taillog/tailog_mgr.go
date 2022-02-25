package taillog

import (
	"fmt"
	"time"
	"go_dev/day12/logAgent2/etcd"
)
/*global variable can be use*/
var taskMgr *Taillogmgr
/*Taillogmgr Managing tasks;
  can be compare new or old configuration
  //LogEntry collect log information
	type LogEntry struct{
		Path string `json:"path"`
		Topic string `json:"topic"`
	}
  control tasks Map 
  type Tailtask struct{
	path string
	topic string
	instance *tail.Tail
}
new ConfChan comparison old ConfChan this is a No buffer area channel
*/
type Taillogmgr struct{
   LogEntry []*etcd.LogEntry 
   taskMap map[string]*Tailtask  
   newConfChan chan []*etcd.LogEntry 
}
/*Init init tail log manager to make  tailTasks 
  new a Taillogmgr 
  range the LogEntry for NewTailtask(path,topic string)
  init it start how many tailtask for judge 
*/
func Init(logConf []*etcd.LogEntry) {
	taskMgr = &Taillogmgr{
		LogEntry:logConf,
		taskMap:make(map[string]*Tailtask,16),
		newConfChan:make(chan []*etcd.LogEntry),
	}
	for _, t := range taskMgr.LogEntry {
		tailObj :=NewTailtask(t.Path,t.Topic)
		mk :=fmt.Sprintf("%s_%s",t.Path,t.Topic)
		taskMgr.taskMap[mk]=tailObj
	}
	go taskMgr.run()
}
/*watch self newConfChan have There are new configurations to do relative processing
  1.add config if no change continue else add config
  2.del config
  3.update config
*/
func (t *Taillogmgr)run(){
	for{
		select {
			case newConf:=<-t.newConfChan:
				for _, conf := range newConf {
					mk :=fmt.Sprintf("%s_%s",conf.Path,conf.Topic)
					_,ok :=t.taskMap[mk]
					if ok {
							continue
					}else{
						tailObj:=NewTailtask(conf.Path,conf.Topic)
						t.taskMap[mk]=tailObj
					}
				}

				for _, old := range t.LogEntry {
					del :=false
					for _, new := range newConf{
						if old.Path==new.Path&&old.Topic==new.Topic{
							continue
						}
					}
					if del {
						mk :=fmt.Sprintf("%s_%s",old.Path,old.Topic)
						t.taskMap[mk].cancelFunc()
					}
				}

				fmt.Println("新的配置来啦",newConf)
			default:
				time.Sleep(time.Second)
		}
	}
}
/*NewConfChan out tskMgr.newConfChan only write*/
func NewConfChan()chan<- []*etcd.LogEntry  {
	return taskMgr.newConfChan
}