package main

import (
	"context"
	"fmt"
	"time"
	"go.etcd.io/etcd/clientv3"
)

func main() {
	client,err :=clientv3.New(clientv3.Config{//创建一个客户端
		Endpoints:[]string{"localhost:2379","localhost22379","localhost:323379"},
		DialTimeout: 5 * time.Second,//创建超时时间
	})
	if err!=nil{//如果err错了
		fmt.Println("connect failed,err:",err)
		return
	}
	fmt.Println("connect succ")
	defer client.Close()//结束的时候关闭客户端
	client.Put(context.Background(),"logagent/conf/","1231231")//往etcd推送消息
	for {//不断地去读取
		rch :=client.Watch(context.Background(),"logagent/conf/")//监听这个对象的
		for wresp:= range rch {//因为上下文对象
			for _, ev := range wresp.Events {
				fmt.Printf("%s %q : %q\n",ev.Type,ev.Kv.Key,ev.Kv.Value)
			}
		}
	}
}