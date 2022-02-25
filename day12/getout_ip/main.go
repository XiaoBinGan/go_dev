package main

import (
	"strings"
	"fmt"
	"net"
)

/*GetOutboundIp getip*/
func GetOutboundIp()(ip string,err error){
	conn,err :=net.Dial("udp","8.8.8.8")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer conn.Close()
	localAddr :=conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.IP.String(),":")[0]
	return
}

func main() {
	ip ,err :=GetOutboundIp()
	if err!=nil{
		// fmt.Println(err)
	}
	fmt.Println(ip)
}