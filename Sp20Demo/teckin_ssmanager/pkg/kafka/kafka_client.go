package kafka

import (
	"context"
	"fmt"
	"log"
	"strings"
	"teckin_ssmanager/config"

	"github.com/Shopify/sarama"
)



type Kafka struct {
	config       *config.KafKaConfig
	responseFunc func(string)
}

func NewKafka(conf *config.Config) *Kafka {
	return &Kafka{config: conf.Kafka, responseFunc: nil}
}

/*
消费队列
topics 主题，切片
responseFunc 回调函数
*/
func (p *Kafka) Consumer(topics []string, responseFunc func(string)) {
	version, err := sarama.ParseKafkaVersion(p.config.Version)
	if err != nil {
		fmt.Println("Error parsing Kafka version:", err)
	}
	config := sarama.NewConfig()
	config.Version = version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange // 分区分配策略
	config.Consumer.Offsets.Initial = -2                                   // 未找到组消费位移的时候从哪边开始消费
	config.ChannelBufferSize = p.config.ChannelBufferSize                  // channel长度

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(p.config.Brokers, ";"), p.config.Group, config)
	if err != nil {
		fmt.Println("Error creating consumer group client:", err)
	}
	p.responseFunc = responseFunc
	go func() {
		for {
			if err := client.Consume(ctx, topics, p); err != nil {
				fmt.Println("Error from consumer:", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				log.Println(ctx.Err())
				// 失败则一直重试
				cancel()
				if err = client.Close(); err != nil {
					fmt.Println("Error closing client:", err)
				}
				p.Consumer(topics, responseFunc)
				return
			}
			p.config.Ready = make(chan bool)
		}
	}()
	<-p.config.Ready
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (p *Kafka) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(p.config.Ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (p *Kafka) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (p *Kafka) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	// 具体消费消息
	for message := range claim.Messages() {
		msg := string(message.Value)
		p.responseFunc(msg)
		//log.Infof("msg: %s", msg)
		//time.Sleep(time.Second)
		//run.Run(msg)
		// 更新位移
		session.MarkMessage(message, "")
	}
	return nil
}

/*
初始化
*/

/*
发送异步消息
topic 主题 string
messages 多个消息。切片类型
*/
func (k *Kafka) Producer(topic string, messages []string) {

	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	version, _ := sarama.ParseKafkaVersion(k.config.Version)
	config.Version = version

	//使用配置,新建一个异步生产者
	producer, e := sarama.NewAsyncProducer(strings.Split(k.config.Brokers, ";"), config)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer producer.AsyncClose()

	//循环判断哪个通道发送过来数据.
	go func(p sarama.AsyncProducer) {
		for {
			select {
			case <-p.Successes():
				return
			case fail := <-p.Errors():
				fmt.Println("kafka err: ", fail.Err)
				return
				//default:
				//	runtime.Goexit()
			}
		}
	}(producer)
	// 发送的消息,主题。
	for _, message := range messages {
		// 注意：这里的msg必须得是新构建的变量，不然你会发现发送过去的消息内容都是一样的，因为批次发送消息的关系。
		msg := &sarama.ProducerMessage{
			Topic: topic,
		}
		//将字符串转化为字节数组
		msg.Value = sarama.ByteEncoder(message)
		//fmt.Println(value)

		//使用通道发送
		producer.Input() <- msg
	}

}
