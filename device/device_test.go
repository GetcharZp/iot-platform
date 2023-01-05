package main

import (
	"context"
	"fmt"
	"gitee/getcharzp/iot-platform/device/device"
	"gitee/getcharzp/iot-platform/device/internal/config"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"testing"
)

var deviceClient device.Device

func init() {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	deviceClient = device.NewDevice(zrpc.MustNewClient(c.RpcClientConf))
}

func TestSendMessage(t *testing.T) {
	reply, err := deviceClient.SendMessage(context.Background(), &device.SendMessageRequest{
		ProductKey: "1",
		DeviceKey:  "device_key",
		Data:       "hello world",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", reply)
}
