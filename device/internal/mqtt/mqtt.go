package mqtt

import (
	"fmt"
	"gitee/getcharzp/iot-platform/models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"strings"
)

var topic = "/sys/#"
var MC mqtt.Client

func NewMqttServer(mqttBroker, clientID, password string) {
	opt := mqtt.NewClientOptions().AddBroker(mqttBroker).SetClientID(clientID).
		SetUsername("get").SetPassword(password)

	// 回调
	opt.SetDefaultPublishHandler(publishHandler)

	MC = mqtt.NewClient(opt)

	// 连接
	if token := MC.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅主题
	if token := MC.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	defer func() {
		// 取消订阅
		if token := MC.Unsubscribe(topic); token.Wait() && token.Error() != nil {
			log.Println("[ERROR] : ", token.Error())
		}
		// 关闭连接
		MC.Disconnect(250)
	}()

	select {}
}

func publishHandler(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("MESSAGE : %s\n", message.Payload())
	fmt.Printf("TOPIC : %s\n", message.Topic())

	topicArray := strings.Split(strings.TrimPrefix(message.Topic(), "/"), "/")
	if len(topicArray) >= 4 {
		if topicArray[3] == "ping" {
			err := models.UpdateDeviceOnlineTime(topicArray[1], topicArray[2])
			if err != nil {
				log.Printf("[DB ERROR] : %v\n", err)
			}
		}
	}
}
