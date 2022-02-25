package main

import (
	"sync"
	// "time"
	"fmt"
	"strings"
	"github.com/Shopify/sarama"
)
var (
	wg sync.WaitGroup
)
func main() {
	//注册一个消费者(kafka)func NewConsumer(addrs []string, config *Config) (Consumer, error) 
	//consumer,err :=sarama.NewConsumer(strings.Split("127.0.0.1:9092",","),nil)
	consumer,err :=sarama.NewConsumer(strings.Split("100.100.69.100:9092",","),nil)
	if err!=nil{
		fmt.Printf("Failed to start consumer: %s",err)
		return
	}
	//partitionList,err :=consumer.Partitions("my_log")//读取topic分区的数量然后返回。
	partitionList,err :=consumer.Partitions("BOX.EVENT_NOTIFIER")//读取topic分区的数量然后返回。
	if err!=nil{
		fmt.Println("Failed to get the list of partitions:",err)
		return
	}
	fmt.Println(partitionList)
	//range 分区  每个分区起一个goroute取消费 

	for partition := range partitionList {
		//消费分区
		//pc,err :=consumer.ConsumePartition("my_log",int32(partition),sarama.OffsetNewest) //ConsumePartition(topic string, partition int32, offset int64) (PartitionConsumer, error)
		pc,err :=consumer.ConsumePartition("BOX.EVENT_NOTIFIER",int32(partition),sarama.OffsetNewest) //ConsumePartition(topic string, partition int32, offset int64) (PartitionConsumer, error)
		if err!=nil{
			fmt.Printf("Failed to start consumer for partion %d %s\n",partition,err)
			return
		}
		defer pc.AsyncClose()//如果报错的话就关闭客户端
		go func (sarama.PartitionConsumer)  {//传入的消费者对象
			wg.Add(1)//等待数量+1
			for msg :=range pc.Messages(){//返回一个channl然后去遍历读取
				fmt.Printf("Partions:%d ,Offset:%d,Key:%s,Value:%s\n",msg.Partition,msg.Offset,string(msg.Key),string(msg.Value))
			}	
			wg.Done()//等待数量-1
		}(pc)//闭包传入
	}
	wg.Wait()//当等待的数量等于0的时候Wait函数返回
	// time.Sleep(time.Hour)
	consumer.Close()//消费结束之后关闭当前的分区

}

// // Consumer manages PartitionConsumers which process Kafka messages from brokers. You MUST call Close()
// // on a consumer to avoid leaks, it will not be garbage-collected automatically when it passes out of
// // scope.
// type Consumer interface {
// 	// Topics returns the set of available topics as retrieved from the cluster
// 	// metadata. This method is the same as Client.Topics(), and is provided for
// 	// convenience.
// 	Topics() ([]string, error)

// 	// Partitions returns the sorted list of all partition IDs for the given topic.
// 	// This method is the same as Client.Partitions(), and is provided for convenience.
// 	Partitions(topic string) ([]int32, error)

// 	// ConsumePartition creates a PartitionConsumer on the given topic/partition with
// 	// the given offset. It will return an error if this Consumer is already consuming
// 	// on the given topic/partition. Offset can be a literal offset, or OffsetNewest
// 	// or OffsetOldest
// 	ConsumePartition(topic string, partition int32, offset int64) (PartitionConsumer, error)

// 	// HighWaterMarks returns the current high water marks for each topic and partition.
// 	// Consistency between partitions is not guaranteed since high water marks are updated separately.
// 	HighWaterMarks() map[string]map[int32]int64

// 	// Close shuts down the consumer. It must be called after all child
// 	// PartitionConsumers have already been closed.
// 	Close() error
// }