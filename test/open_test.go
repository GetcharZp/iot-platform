package test

import (
	"encoding/json"
	"fmt"
	"gitee/getcharzp/iot-platform/helper"
	"testing"
)

var openServiceAddr = "http://127.0.0.1:16001"

func TestSendMessage(t *testing.T) {
	data, _ := json.Marshal(map[string]interface{}{
		"product_key": "1",
		"device_key":  "device_key",
		"data":        "hello world",
	})
	rep, err := helper.HttpPost(openServiceAddr+"/send-message", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
