package main

import (
	"fmt"
	"net"
)

func main() {
	/*开启监听一个客户端
		1.申明一个可以操作的udp服务端
		2.容错
		3.在服务端申明成功的时候申明一个defer 在函数结束的额时候释放服务端端所占的内存
	*/
	udpconn,err :=net.ListenUDP("udp",&net.UDPAddr{
		IP:net.IPv4(0,0,0,0),
		Port:3000,
	})
	if err!=nil{
		fmt.Println("listen failed,err:",err)
		return
	}
	defer udpconn.Close()//上面没有报错直接关闭因为上面可能直接报错可能会没有udpconn
	/*循环的的可以发送请求
		1.创建一个字节数组去接收请求时需要发送的参数
		2.接收参数
		3.容错防止读取数据失败
		4.发送数据到udp客户端
		5.容错防止数据发送失败
	*/
	for{
		var data [1024]byte
		n,udpaddr,err :=udpconn.ReadFromUDP(data[:])//接收参数 和返回发送数据的udp客户端的配置地址和端口
		if err!=nil{
			fmt.Println("read upd failed,err:",err)
			continue
		}
		fmt.Printf("data:%v,addr:%v,count:%v\n",string(data[:n]),udpaddr,n)
		_,err =udpconn.WriteToUDP(data[:n],udpaddr)//发送用户自己传来的数据 到指定的udp客户端
		if err!=nil{
			fmt.Println("write to udp failed,err:",err)
			continue
		}
	}

	 
	 
}