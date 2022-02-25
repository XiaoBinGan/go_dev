package kafka

import (
	"github.com/astaxie/beego/logs"
	"github.com/Shopify/sarama"
)
var (
	client sarama.SyncProducer//创建一个参数接收初始化返回的操作对象
)
//初始化kafka配置
//string addr
func InitKafka(addr string)(err error) {
	config :=sarama.NewConfig() //获取一个kafka对象
	config.Producer.RequiredAcks =sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	
	client,err =sarama.NewSyncProducer([]string{addr},config)// NewSyncProducer使用给定的代理地址和配置创建新的SyncProducer。
	if err!=nil{
		logs.Error("init kafka producer failed, err:",err)
		return
	}
	
	logs.Debug("init kafka succ ")
	return
    
}
//发送消息到kafka
func SendToKafka(data,topic string)(err error)  {
	defer client.Close()
	//创建一个kafka消费对象
	msg :=&sarama.ProducerMessage{}
	msg.Topic = topic //消息发往哪个topic
	msg.Value=sarama.StringEncoder(data)//将数据编码存入kafka

	// partition,offset,err :=client.SendMessage(msg)
	_,_,err =client.SendMessage(msg)
	if err!=nil{
		logs.Error("send massage failed,err:%v data:%v topic:%v",err,data,topic)
		return
	}

	// logs.Debug("send succ,pid:%v offset:%v, topic:%v\n",partition,offset,topic)
	return
}





// SyncProducer publishes Kafka messages, blocking until they have been acknowledged. It routes messages to the correct
// broker, refreshing metadata as appropriate, and parses responses for errors. You must call Close() on a producer
// to avoid leaks, it may not be garbage-collected automatically when it passes out of scope.
//
// The SyncProducer comes with two caveats: it will generally be less efficient than the AsyncProducer, and the actual
// durability guarantee provided when a message is acknowledged depend on the configured value of `Producer.RequiredAcks`.
// There are configurations where a message acknowledged by the SyncProducer can still sometimes be lost.
//
// For implementation reasons, the SyncProducer requires `Producer.Return.Errors` and `Producer.Return.Successes` to
// be set to true in its configuration.
// type SyncProducer interface {

	// SendMessage produces a given message, and returns only when it either has
	// succeeded or failed to produce. It will return the partition and the offset
	// of the produced message, or an error if the message failed to produce.
	// SendMessage(msg *ProducerMessage) (partition int32, offset int64, err error)

	// SendMessages produces a given set of messages, and returns only when all
	// messages in the set have either succeeded or failed. Note that messages
	// can succeed and fail individually; if some succeed and some fail,
	// SendMessages will return an error.
	// SendMessages(msgs []*ProducerMessage) error

	// Close shuts down the producer and waits for any buffered messages to be
	// flushed. You must call this function before a producer object passes out of
	// scope, as it may otherwise leak memory. You must call this before calling
	// Close on the underlying client.
	// Close() error
// }

// make strings and byte slices encodable for convenience so they can be used as keys
// and/or values in kafka messages

// StringEncoder implements the Encoder interface for Go strings so that they can be used
// as the Key or Value in a ProducerMessage.
// msg.value=sarama.StringEncoder(data)
