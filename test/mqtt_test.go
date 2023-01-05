package test

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"testing"
	"time"
)

func TestMqtt(t *testing.T) {
	opt := mqtt.NewClientOptions().AddBroker("tcp://192.168.1.8:1883").SetClientID("go-test").
		SetUsername("get").SetPassword("123456")

	// 回调
	opt.SetDefaultPublishHandler(func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("MESSAGE : %s\n", message.Payload())
		fmt.Printf("TOPIC : %s\n", message.Topic())
	})

	c := mqtt.NewClient(opt)

	// 连接
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}

	// 订阅主题
	if token := c.Subscribe("/sys/1/device_key/receive", 0, nil); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}

	// 发布
	if token := c.Publish("/sys/1/device_key/ping", 0, false, "Hello"); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}

	time.Sleep(time.Second * 3600 * 24)

	// 取消订阅
	if token := c.Unsubscribe("/topic/#"); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}
	// 关闭连接
	c.Disconnect(250)
}
