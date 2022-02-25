package main

import (
	"go.etcd.io/etcd/mvcc/mvccpb"
	"encoding/json"
	"go_dev/day11/logCollect/tailf"
	"github.com/astaxie/beego/logs"
	"context"
	"strings"
	"fmt"
	"time"
	etcd_client "go.etcd.io/etcd/clientv3"
)


type EtcdClient struct{
	client *etcd_client.Client  //创建类型方便全局使用
	keys  []string		//存放configKey +arr组合的etcdKey
}

var (
	etcdClient *EtcdClient//申明一个全局对象方便操作etcd对象
	// collectConf []tailf.CollectConf //申明一个CollectConf反序列化数据信息
)
//初始化etcd
func initEtcd(addr string,key string)(collectConf []tailf.CollectConf,err error) {
		client,err :=etcd_client.New(etcd_client.Config{//创建ercd的客户端
			Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
			// Endpoints: []string{"192.168.1.3:2379"},
			DialTimeout: 5* time.Second,
		})
		if err != nil {
			fmt.Println("connect failed, err:", err)
			return
		}
		etcdClient = &EtcdClient{//初始化一个全局对象接收etcd的客户端
			client:client,
		}
		//如果发现配置文件没有已/后缀结尾就加上防止报错
		if strings.HasSuffix(key,"/")==false{
			key = key+"/"
		}
		//遍历所有获取到的网卡ip
		for _, ip := range localIpArray {
			etcdKey :=fmt.Sprintf("%s%s",key,ip)
			etcdClient.keys=append(etcdClient.keys,etcdKey)
			fmt.Println("etcdKey:",etcdKey)
			ctx,cancel :=context.WithTimeout(context.Background(),time.Second)//注册一个上下文超时
			resp,err :=client.Get(ctx,etcdKey)//得到检索键。
			if err!=nil{
				logs.Error("client get from etcd failed,err:%v",err)
				continue//如果失败了就继续获取
			}
			cancel()//关闭当前的上下文
			logs.Debug("resp from etcd:%v",resp.Kvs)
			//如果Get到了结果那么遍历返回结果中Kvs
			for _, v := range resp.Kvs{//kvs is the list of key-value pairs matched by the range request.
				if  string(v.Key)==etcdKey{
					err = json.Unmarshal(v.Value,&collectConf)//将符合条件的参数返序列化到collectConf中
					
					continue
				}
				logs.Debug("log config is %v",collectConf)
				// fmt.Println(k,v)
			}
		}
		initEtcdWatcher()
		return
}
func initEtcdWatcher()  {
	//遍历存放configKey +arr组合的etcdKey
	for _, key := range etcdClient.keys {
		go watchKey(key)
	}
}
func watchKey(key string)  {//监听etcd 当中key
	client,err :=etcd_client.New(etcd_client.Config{
		Endpoints :[]string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5*time.Second,
	})
	if err!=nil{
		logs.Error("connect etcd failed ,err:",err)
		return
	}
	logs.Debug("bengin watch key:%s",key)
	for {
		rch :=client.Watch(context.Background(),key)//将当前的key传入上下文
		var collectConf []tailf.CollectConf			//创建一个collectConf来接收logpath 和topic
		var getConfSucc = true						//如果获取成功的的标志位
		for wresp := range rch {					//遍历上下文中的每一项
			for _, ev := range wresp.Events {	    //返回一个channl然后遍历里面每一项
				if ev.Type == mvccpb.DELETE{		//判断类型是否是删除
					logs.Warn("key[%s] 's config deleted",key)
					continue
				}
				if ev.Type==mvccpb.PUT&&string(ev.Kv.Key)==key{			//判断类型是否是添加
					err = json.Unmarshal(ev.Kv.Value,&collectConf)		//如果是添加就反序列化参数到collectConf []tailf.CollectConf当中
					if err!=nil{
						logs.Error("key [%s],Unmarshal[%s],err:%v",err)
						getConfSucc=false								//如果反序列化失败了没有获取到参数配置
						continue
					}
				}
				logs.Debug("get config ftom etcg, %s %q %q\n", ev.Type,ev.Kv.Key,ev.Kv.Value)
			}
			if getConfSucc{												//如果标志位是获取配置成功那么自动更新配置
				logs.Debug("get config from etcd succ,%v",collectConf)	//如果这里可以打印日志说明这里没问题那么问题就出现早tailf文件当中
				tailf.UpdateConfig(collectConf)							//将添加的内容传入函数进行追踪
			}
		}
	}
}