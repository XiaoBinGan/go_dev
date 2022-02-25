package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var (
	//client needed in many places use 
	client sarama.SyncProducer
)

/*Init inti kafka connect
	params {addrs []string{"192.163.0.1:9092","192.163.0.2:9092"}}
	new Saramconfig 
	configAcks send end need tell with leader and follow
	Partitioner chose an new pration
	Success  send massage well be in success channl return
*/
func Init(addrs []string)(err error) {
	config :=sarama.NewConfig()
	config.Producer.RequiredAcks=sarama.WaitForAll
	config.Producer.Partitioner=sarama.NewRandomPartitioner
	config.Producer.Return.Successes=true
	client,err =sarama.NewSyncProducer(addrs,config)
	if err!=nil{
		fmt.Printf("client connect failed,err:%v\n",err)
		return err
	}
	return nil
}
/*SendToKafka send massage to kafka's topic
	params{
		topic string   save massage chanl
		lineMsg string log
	}
	construct a massage 
	spacifies the theme of the store
	use client.SendMessage(massage)
*/
func SendToKafka(topic,lineMsg string)  {
	msg:=&sarama.ProducerMessage{}
	msg.Topic=topic
	msg.Value=sarama.StringEncoder(lineMsg)
	partition, offset,err :=client.SendMessage(msg)
	if err!=nil{
		fmt.Printf("send massage to kafa topic:%v failed,err%v\n ", topic,err)
		return
	}
	fmt.Printf("send massage success to topic:%v partion:%v offset:%v\n",topic,partition,offset)
}
