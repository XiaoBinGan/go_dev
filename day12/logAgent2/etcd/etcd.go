package etcd

import (
	"encoding/json"
	"context"
	"fmt"
	"time"
	etcd_client "go.etcd.io/etcd/clientv3"
)

var (
	/*The variable for global use*/
	client *etcd_client.Client
)

//LogEntry need collect log information
type LogEntry struct{
	Path string `json:"path"`
	Topic string `json:"topic"`
}
/*Init etcd config 
	params{
		Endpoints:[]string{"localhost:2379"},
		DialTimeout:time.Duration*time.second,
	}
*/
func Init(addrs []string,timeout time.Duration)(err error)  {
	client,err= etcd_client.New(etcd_client.Config{
		Endpoints:addrs,
		DialTimeout:timeout,
	})
	if err!=nil{
		fmt.Printf("etcd connect failed ,err:%v\n",err)
		return err
	}
	return nil 
}

/*GetConf get etcd config detail
  params{
	  key /logagent/collect_config
  }
  return (
	 LogEntryConf []*LogEntry  
	 err		   error
  )
  use context.WithTimeout(context.Background,time.Second) to control DialTimeout
  use global variable `client` to receive the client.Get(ctx,key) return 
  use cancel() close context
  cover err 
  for _,ev :=range resp.KvsP{
	json.unmarshal(ev.key,LogEntryConf)  
	ev.key,ev.value}
*/
func GetConf(key string )(LogEntryConf []*LogEntry,err error){
	ctx,cancel :=context.WithTimeout(context.Background(),time.Second)
	resp,err :=client.Get(ctx,key)
	fmt.Printf("resp:%#v\n",resp)
	cancel()
	if err!=nil{
		fmt.Printf("get key failed ,err:%v\n",err)
		return
	}
	for _, ev := range resp.Kvs {
		err :=json.Unmarshal(ev.Value,&LogEntryConf)
		if err!=nil{
			fmt.Printf("json unmarshal failed,err:%v\n",err)
			return nil,err
		}
	}
	return 
}
/*WatchConf Monitor etcd by key configuration 
  watch client.Watch(context.Background(),key)
  range the watch reslut watchChan
  if ev.Type!=etcd_client.EventTypeDelete not delete
		  get wresp.Events[ev].Kv.Value unmarshal to the newConfChan
     else
  send newConf to the newConfch(chan<- []*LogEntry only write chanl)
*/
func WatchConf(key string,newConfch chan<- []*LogEntry)  {
	watchChan :=client.Watch(context.Background(),key)
	for wresp:= range watchChan  {
		for _, ev := range wresp.Events {
			fmt.Printf("Type:%v,Key:%s,Value:%v\n",ev.Type,string(ev.Kv.Key),string(ev.Kv.Value))
			 var newConf []*LogEntry
			 if ev.Type!=etcd_client.EventTypeDelete{
				 err :=json.Unmarshal(ev.Kv.Value,&newConf)
				 if err!=nil{
					fmt.Printf("json Unmarshal failed,err:%v\n",err)
					continue
				 }
			 }
			fmt.Printf("get now config success")
            newConfch <-newConf
		}
	}
}