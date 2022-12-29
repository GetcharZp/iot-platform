package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

func NewMqttServer(mqttBroker string) {
	opt := mqtt.NewClientOptions().AddBroker(mqttBroker).SetClientID("go-mqtt-server-client-id")

	// 回调
	opt.SetDefaultPublishHandler(func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("MESSAGE : %s\n", message.Payload())
		fmt.Printf("TOPIC : %s\n", message.Topic())
	})

	c := mqtt.NewClient(opt)

	// 连接
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅主题
	if token := c.Subscribe("/topic/#", 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	defer func() {
		// 取消订阅
		if token := c.Unsubscribe("/topic/#"); token.Wait() && token.Error() != nil {
			log.Println("[ERROR] : ", token.Error())
		}
		// 关闭连接
		c.Disconnect(250)
	}()

	select {}
}
