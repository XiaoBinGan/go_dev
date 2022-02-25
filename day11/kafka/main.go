package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll //发送完数据需要跟leader和follow确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner//新选出一个partition
	config.Producer.Return.Successes = true//成功交付的消息将在success channl中返回

	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)//连接卡夫卡客户端
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}

	defer client.Close()//结束的时候关闭连接
	//制造一条记录
	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"//
	msg.Value = sarama.StringEncoder("this is a good test, my message is good")
	//向客户端发送完成后记录下的偏移量和位置
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}

	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
