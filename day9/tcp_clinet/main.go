package main

import (
	"strings"
	"os"
	"bufio"
	"fmt"
	"net"
)

func Send()  {
		// 	Dial是应用程序提供的用于创建和配置a的功能
		// 连接。
		// 从拨号返回的连接不能处于特殊状态
		// (订阅pubsub频道，交易开始，…)
	conn,err :=net.Dial("tcp","localhost:50000")
	if err!=nil {
		fmt.Println("dialing failed err:",err.Error())
	}
	defer conn.Close()
	inputReader :=bufio.NewReader(os.Stdin)		//创建缓冲区 读取控制台输入的文件
	for {
		data,_ :=inputReader.ReadString('\n')   //读取文件以换行结尾
		tirmmedInput :=strings.Trim(data,"\r\n")//去除Linux 的换行Windows换行
		if tirmmedInput=="Q" {
			fmt.Println("application exit")
			return
		}
		count,err :=conn.Write([]byte(tirmmedInput)) //count输入要素的长度
		fmt.Println(count)
		if err!=nil {
			fmt.Println("connection failed,err:",err.Error())
			return
		}
	}
}



func main() {
	Send()
}



//使用net的Dial方法新建一个链接
//容错
//函数结束的时候关闭这个链接
// 创建一个bufio去读取命令行的输入要素
//写一个死循环
//读取输入的









// func Dial
// func Dial(network, address string) (Conn, error)
// 在网络network上连接地址address，并返回一个Conn接口。可用的网络类型有：

// "tcp"、"tcp4"、"tcp6"、"udp"、"udp4"、"udp6"、"ip"、"ip4"、"ip6"、"unix"、"unixgram"、"unixpacket"

// 对TCP和UDP网络，地址格式是host:port或[host]:port，参见函数JoinHostPort和SplitHostPort。

// Dial("tcp", "12.34.56.78:80")
// Dial("tcp", "google.com:http")
// Dial("tcp", "[2001:db8::1]:http")
// Dial("tcp", "[fe80::1%lo0]:80")
// 对IP网络，network必须是"ip"、"ip4"、"ip6"后跟冒号和协议号或者协议名，地址必须是IP地址字面值