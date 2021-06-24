package kafka

//import (
//	"fmt"
//	"testing"
//)
//
//func TestMqttClient(t *testing.T) {
//	kafkaConf := &KafKaConfig{
//		Brokers:           "127.0.0.1:9092",
//		Version:           "2.2.1",
//		Group:             "test",
//		ChannelBufferSize: 256, //TODO: 先使用默认设置的值，后期根据需要进行调整
//		Ready:             make(chan bool),
//	}
//	KafkaClient := NewKafka(kafkaConf)
//	KafkaClient.Consumer([]string{"Topic"}, func(kfkMsg string) {
//		fmt.Println(kfkMsg)
//	})
//}
