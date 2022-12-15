package test

import (
	"encoding/json"
	"fmt"
	"gitee/getcharzp/iot-platform/helper"
	"testing"
)

var adminServiceAddr = "http://127.0.0.1:14010"

func TestDeviceList(t *testing.T) {
	m := map[string]string{
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaWRlbnRpdHkiOiIxIiwibmFtZSI6ImdldCIsImV4cCI6MTY3MzUzMTY1N30.XNyNdGwe8xYS7RvPM-LcZph_ade8EfEV2fMV7WPSmZc",
	}
	header, _ := json.Marshal(m)
	rep, err := helper.HttpGet(adminServiceAddr+"/device/list?page=1&size=20&name=", header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestProductList(t *testing.T) {
	m := map[string]string{
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaWRlbnRpdHkiOiIxIiwibmFtZSI6ImdldCIsImV4cCI6MTY3MzUzMTY1N30.XNyNdGwe8xYS7RvPM-LcZph_ade8EfEV2fMV7WPSmZc",
	}
	header, _ := json.Marshal(m)
	rep, err := helper.HttpGet(adminServiceAddr+"/product/list?page=1&size=20&name=", header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
