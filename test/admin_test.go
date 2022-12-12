package test

import (
	"fmt"
	"gitee/getcharzp/iot-platform/helper"
	"testing"
)

var adminServiceAddr = "http://127.0.0.1:14010"

func TestDeviceList(t *testing.T) {
	rep, err := helper.HttpGet(adminServiceAddr + "/device/list?page=1&size=20&name=")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
