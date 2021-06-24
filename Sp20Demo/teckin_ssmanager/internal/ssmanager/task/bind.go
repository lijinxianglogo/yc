package task

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strconv"
	"strings"
	"teckin_ssmanager/internal/ssmanager/model/mqttModel"
	"teckin_ssmanager/internal/ssmanager/service"
)

func InitBind(srv *service.Service) {
	_, err := srv.Dao.EmqxClient.Subscribe("test/device/bind", 0, func(client mqtt.Client, message mqtt.Message) {
		fmt.Println("收到：" + string(message.Payload()))
		var bindInfo mqttModel.ApplianceControlBind
		err := json.Unmarshal(message.Payload(), &bindInfo)
		if err != nil {
			fmt.Println("Bind info unmarshal fail：", err)
		} else {
			fmt.Println("Bind info unmarshal success")
			//取到uuid
			fromSli := strings.Split(bindInfo.Header.From, "/")
			if len(fromSli) < 3 {
				fmt.Println("failed to get uuid from header. from:" + bindInfo.Header.From)
			} else {
				fmt.Println("get uuid success. uuid=", fromSli[2])
				uuid := fromSli[2]
				uid := bindInfo.Payload.Bind.Uid
				fmt.Println("get uid success. uid=", uid)
				//更新数据库
				err := srv.Dao.Mysql.BindDevice(uuid, uid)
				var msgAck mqttModel.ApplianceControlBindAck
				msgAck.Header = bindInfo.Header
				msgAck.Header.From = "test/cloud"
				msgAck.Header.Src = "cloud"
				msgAck.Header.Method = "SETACK"
				if err != nil {
					fmt.Println("uid=", uid)
					fmt.Println("failed to bind device. err:", err)
					msgAck.Payload = mqttModel.BindAckPayload{Result: mqttModel.Result{Status: 2, Uuid: uuid}}

				} else {
					fmt.Println("bind device success")
					msgAck.Payload = mqttModel.BindAckPayload{Result: mqttModel.Result{Status: 1, Uuid: uuid}}
				}

				//发送bindAck
				ackStr, _ := json.Marshal(msgAck)
				_, err = srv.Dao.EmqxClient.Publish(bindInfo.Header.From, string(ackStr))
				if err != nil {
					fmt.Println("publish mqtt msg to device fail", err)
				} else {
					fmt.Println("publish mqtt msg to device success")
				}
				//向APP发送结果
				_, err = srv.Dao.EmqxClient.Publish(mqttModel.AppOnlineStatusTopic+strconv.Itoa(uid), string(ackStr))
				if err != nil {
					fmt.Println("publish mqtt msg to app fail", err)
				} else {
					fmt.Println("publish mqtt msg to app success")
				}
			}
		}
	})
	if err != nil {
		fmt.Println("Subscribe err:", err)
	}
}
