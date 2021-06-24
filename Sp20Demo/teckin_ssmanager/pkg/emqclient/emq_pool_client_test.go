package emqclient

//
//import (
//	"fmt"
//	"mssiot_device/app/deps"
//	"mssiot_device/app/library"
//	"mssiot_device/app/library/handleemqmqtt"
//	"sync"
//	"testing"
//)
//
//func testNewEmqxClient(t *testing.T) {
//	payload := make(map[string]interface{})
//	toggle := make(map[string]interface{})
//	toggle["onoff"] = 0
//	toggle["channel"] = 3
//	payload["togglex"] = toggle
//	from := handleemqmqtt.CreateApplianceSubscribe("1909100843416000040534298f1f1f54")
//	namespace := "Appliance.Control.ToggleX"
//	method := "GET"
//	key := "22b19deb1548a4f89bb9efd7e8ab7ac2"
//	triggerSrc := library.TriggerSrcCloud
//	topic := handleemqmqtt.CreateApplianceSubscribe("1909100843416000040534298f1f1f54")
//	wg := sync.WaitGroup{}
//	wg.Add(10)
//	// 协程使用
//	for i := 0; i < 10; i++ {
//		go func() {
//			message, _ := handleemqmqtt.CreateSendToEmqxMessage(topic, namespace, method, key, payload, triggerSrc)
//			msg, err := deps.EmqxClient.SyncPublish(from, topic, message)
//			fmt.Println(msg, err)
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//
//	// 普通使用
//	for i := 0; i < 10; i++ {
//		message, _ := handleemqmqtt.CreateSendToEmqxMessage(topic, namespace, method, key, payload, triggerSrc)
//		msg, err := deps.EmqxClient.SyncPublish(from, topic, message)
//		fmt.Println(msg, err)
//	}
//}
