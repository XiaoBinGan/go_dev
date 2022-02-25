package kafka

import (
	"time"
	"fmt"
	"github.com/Shopify/sarama"
)
type logData struct{
	Topic string
	Data string
}

var (
	//client needed in many places use 
	client sarama.SyncProducer
	//get taillog send data(topic,line.Text)
	logDataChan chan *logData
)

/*Init inti kafka connect
	params {addrs []string{"192.163.0.1:9092","192.163.0.2:9092"}}
	new Saramconfig 
	configAcks send end need tell with leader and follow
	Partitioner chose an new pration
	make  logDataChan will to 
	create goroutine send real data to kafka xxx topic
	Success  send massage well be in success channl return
*/
func Init(addrs []string,chanMaxSize int)(err error) {
	config :=sarama.NewConfig()
	config.Producer.RequiredAcks=sarama.WaitForAll
	config.Producer.Partitioner=sarama.NewRandomPartitioner
	config.Producer.Return.Successes=true
	client,err =sarama.NewSyncProducer(addrs,config)
	if err!=nil{
		fmt.Printf("client connect failed,err:%v\n",err)
		return err
	}
	logDataChan = make(chan *logData,chanMaxSize)
	go sendToKafka()
	return
}
/*SendToChan get tail log send data to the logDataChan*/
func SendToChan(topic,Data string)  {
	lg:=&logData{
		Topic:topic,
		Data:Data,
	}
	logDataChan<-lg
}
/*SendToKafka send massage to kafka's topic
    range logDataChan every item to kafka topic
	construct a massage 
	spacifies the theme of the store
	use client.SendMessage(massage)
*/
func sendToKafka()  {
	for {
		select {
		case data:=<-logDataChan:
			msg:=&sarama.ProducerMessage{}
			msg.Topic=data.Topic
			msg.Value=sarama.StringEncoder(data.Data)
			partition, offset,err :=client.SendMessage(msg)
			if err!=nil{
				fmt.Printf("send massage to kafa topic:%v failed,err%v\n ", data.Topic,err)
				return
			}
			fmt.Printf("send massage success to topic:%v partion:%v offset:%v\n",data.Topic,partition,offset)
		default:
			time.Sleep(time.Nanosecond)
		}

	}
	
}
