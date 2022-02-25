package main

import (
	"encoding/json"
	"go_dev/day11/logCollect/tailf"
	"context"
	"fmt"
	"time"
	etcd_clinet "go.etcd.io/etcd/clientv3"
)
//因为解析出来的参数需要tailf去跟踪所有存在一些问题 所以直接使用tailf当中的CollectConf
// type LogConf struct{
// 	Path string `json:"path"`
// 	Topic string `json:"topic"`
// }
const(
	EtcdKey = "/logagent/config/192.168.1.3"
	// EtcdKey = "/logagent/config/172.25.1.149"
)
var(
	logConfArr []tailf.CollectConf
)

func SetLogConfToetcd()  {
	client,err :=etcd_clinet.New(etcd_clinet.Config{
		// Endpoints :[]string{"192.168.1.3:2379"},
		Endpoints :[]string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err!=nil{
		fmt.Println("connect failed ,err:",err)
		return
	}
	fmt.Println("connect succ")
	defer client.Close()
	//创建一个超时时间传入上下文操作
	
		logConfArr = append(//添加配置项
			logConfArr,
			tailf.CollectConf{
				LogPath:"/Users/wujiahao/study&demo/GO/package/src/go_dev/day11/logCollect/logs/logagent.log",
				Topic:"nginx_log",
			},
		)
		logConfArr = append(//添加配置项
			logConfArr,
			tailf.CollectConf{
				LogPath:"/Users/wujiahao/study&demo/GO/package/src/go_dev/day11/logCollect/logs/nginx.log",
				Topic:"nginx_log",
			},
		)
	
		logConfArr = append(//添加配置项
			logConfArr,
			tailf.CollectConf{
				LogPath:"/Users/wujiahao/study&demo/GO/package/src/go_dev/day11/logCollect/logs/nginx_err.log",
				Topic:"nginx_log",
			},
		)
	data,err :=json.Marshal(logConfArr)//将配置变成json
	if err !=nil{
		fmt.Println("json failed,",err)
		return
	}

	ctx,cancel :=context.WithTimeout(context.Background(),time.Second)
	// client.Delete(ctx,EtcdKey)
	// return

	_,err= client.Put(ctx,EtcdKey,string(data))//将数据已key-vaule的方式存储在etcd当中
	cancel()//使用完了关闭内容
	if err!=nil{
		fmt.Println("put failed ,err:",err)
		return
	}		

	ctx,cancel = context.WithTimeout(context.Background(),time.Second)
	resp,err :=client.Get(ctx,EtcdKey)//获取etcd当中存储的key的值
	cancel()//关闭上下文的通道
	if err!=nil{
		fmt.Println("get failed ,err:",err)
		return
	}
	for _, ev := range resp.Kvs {//获取到的可能是kv的数组遍历数组的每一项
		fmt.Printf("%s : %s\n",ev.Key,ev.Value)
	} 
}
func  main()  {
	SetLogConfToetcd()
}