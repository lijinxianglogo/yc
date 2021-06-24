## kafka

使用该客户端，连接kafka消费或者生产消息。<br/>
（*注意：例子中所提供的参数为非真模拟参数*）

### 取消息的逻辑
客户端会取出一条消息后直接处理，然后等待处理完之后在取吓一条消息


## 1.首先需要实例化该客户端。

```golang
	kafkaConf := &KafKaConfig{
    		Brokers:           "127.0.0.1:9092",
    		Version:           "2.2.1",
    		Group:             "test",
    		ChannelBufferSize: 256, //TODO: 先使用默认设置的值，后期根据需要进行调整
    		Ready:             make(chan bool),
    	}
    	KafkaClient := NewKafka(kafkaConf)
```
传入参数：
- brokers:kafka集群地址,多个地址使用分号间隔
- version: kafka版本
- group: 客户端消费组名称
- channelBufferSize: 通道大小 (从kafka获取缓存在缓存队列中的数据大小， 处理的时候还是一条一条处理)
- ready: 通道是否正在运行

返回参数：

- client：实例化的客户端


## 2.生产数据
```golang
topic := "test1-topic"
	var message string
	for i:=0;;i++ {
		time.Sleep(500 * time.Millisecond)
		time11 := time.Now()
		message = "this is a message 0606 " + time11.Format("15:04:05")
		kafkaClient.Producer(topic, []string{message})
```
请求参数：
- topic:消息的topic，主题
- message:消息内容。

返回参数：
异步写入，无返回值

## 3.消费数据
```golang
funcA := func(message string) {
		fmt.Println(message)
	}
	topics := []string{"test1-topic"}
	kafkaClient.Consumer(topics, funcA)
```
请求参数：
- topics:消息的topic，主题
- funcA:回调函数。

返回参数：
异步消费，无返回值



