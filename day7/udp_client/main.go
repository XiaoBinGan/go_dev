package main

import (
	"fmt"
	"net"
)

/*UDP客户端
  1.生成upd客户端指定请求的服务端的接口
  2.容错防止客户端生成失败
  3.申明一个defer在函数返回的时候关闭upd客户端释放内存
  4.创建将要发送到服务端的参数字节数组
  5.发送数据
  6.容错防止数据发送失败
  7.创建一个字节数组接收
  8.接收服务端返回的数据
  9.容错如果失败了就直接报错
  10.打印接受的内容

*/
func main() {
	udpClient,err :=net.DialUDP("udp",nil,&net.UDPAddr{
		IP: net.IPv4(0,0,0,0),
		Port:3000,
	})
	if err!=nil{
		fmt.Println("连接服务器失败,err:",err)
		return
	}
	defer udpClient.Close()
	for{
		var msg string
		fmt.Scanln(&msg)
		sendData :=[]byte(msg)//将控制台输入的内容转换成字节切片
		_,err =udpClient.Write(sendData)//字节切片写入请求
		if err!=nil{
			fmt.Println("send msg failed ,err:",err)
			return
		}
		data :=make([]byte, 4096)
		n,UDPAddr,err :=udpClient.ReadFromUDP(data)//将接收的数据写入data字节切片当中
		if err!=nil{
			fmt.Println("接收数据失败 ,err:",err)
			return
		}
		fmt.Printf("recv:%v,addr:%v,count:%v\n",string(data[:n]),UDPAddr,n)	
	}

}