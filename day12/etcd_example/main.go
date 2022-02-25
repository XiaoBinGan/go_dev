package main

import (
	"context"
	"fmt"
	"time"
	etcd_clinet "go.etcd.io/etcd/clientv3"
)
//初始化etcd客户端
//向客户端传入上下文然后Put数据key value
//读取etcd客户端的内容因为读取的内容可能是数组所以
func EtcdExample() {
	client,err :=etcd_clinet.New(etcd_clinet.Config{
		Endpoints :[]string{"localhost:2379","localhost:22379","localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err!=nil{
		fmt.Println("connect failed ,err:",err)
		return
	}
	fmt.Println("connect succ")
	defer client.Close()
	//创建一个超时时间传入上下文操作
	ctx,cancel :=context.WithTimeout(context.Background(),time.Second)
	// _,err= client.Put(ctx,".logagent/conf/","sample_value")//将数据已key-vaule的方式存储在etcd当中¸
	key :=`[{"path":"d:/tmp/nginx","topic":"my_nginx"},{"path":"d:/tmp/nginx","topic":"my_nginx"}]`
	// key :=`[{"path":"d:/tmp/nginx","topic":"my_nginx"},{"path":"d:/tmp/nginx","topic":"my_nginx"},{"path":"f:/tmp/nginx","topic":"my_nginx"}]`
	_,err= client.Put(ctx,"/logagent/collect_config",key)//将数据已key-vaule的方式存储在etcd当中¸
	cancel()//使用完了关闭内容
	if err!=nil{
		fmt.Println("put failed ,err:",err)
		return
	}
	ctx,cancel = context.WithTimeout(context.Background(),time.Second)
	resp,err :=client.Get(ctx,"logagent/conf/")//获取etcd当中存储的key的值
	cancel()//关闭上下文的通道
	if err!=nil{
		fmt.Println("get failed ,err:",err)
		return
	}
	for _, ev := range resp.Kvs {//获取到的可能是kv的数组遍历数组的每一项
		fmt.Printf("%s : %s\n",ev.Key,ev.Value)
	} 

}


type LogConf struct{
	Path string `json:"path"`
	Topic string `json:"topic"`
}
const(
	EtcdKey = "/logagent/conf/127.0.0.1"
)
var(
	logConfArr []LogConf
)

func SetLogConfToetcd()  {
	client,err :=etcd_clinet.New(etcd_clinet.Config{
		Endpoints :[]string{"localhost:2379","localhost:22379","localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err!=nil{
		fmt.Println("connect failed ,err:",err)
		return
	}
	fmt.Println("connect succ")
	defer client.Close()
	//创建一个超时时间传入上下文操作
	ctx,cancel :=context.WithTimeout(context.Background(),time.Second)
	_,err= client.Put(ctx,".logagent/conf/","sample_value")//将数据已key-vaule的方式存储在etcd当中
	cancel()//使用完了关闭内容
	if err!=nil{
		fmt.Println("put failed ,err:",err)
		return
	}
	ctx,cancel = context.WithTimeout(context.Background(),time.Second)
	resp,err :=client.Get(ctx,"logagent/conf/")//获取etcd当中存储的key的值
	cancel()//关闭上下文的通道
	if err!=nil{
		fmt.Println("get failed ,err:",err)
		return
	}
	fmt.Println(resp)
	for _, ev := range resp.Kvs {//获取到的可能是kv的数组遍历数组的每一项
		fmt.Printf("%s : %s\n",ev.Key,ev.Value)
	} 
}
func  main()  {
	// SetLogConfToetcd()
	EtcdExample()
}