package main

import (
	"fmt"
	"net"
)


var (
	localIpArray []string
)



func init() {
	addrs,err :=net.InterfaceAddrs()//获取网卡地址
	if err!=nil {
		// 报错的话直接打印错误信息
		panic(fmt.Sprintf("get local ip failed, %v",err))
	}
	for _, addr := range addrs {//遍历所有的地址
		if ipent,ok :=addr.(*net.IPNet);ok && !ipent.IP.IsLoopback(){
			if ipent.IP.To4()!=nil{
				localIpArray=append(localIpArray,ipent.IP.String())//将所以的ipv4的的网络存储在数组当中
			}
		}
	}
	fmt.Println("localIpArray:",localIpArray)
}