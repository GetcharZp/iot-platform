package test

import (
	"encoding/json"
	"fmt"
	"gitee/getcharzp/iot-platform/helper"
	"testing"
)

var adminServiceAddr = "http://127.0.0.1:14010"
var m1 = map[string]string{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaWRlbnRpdHkiOiIxIiwibmFtZSI6ImdldCIsImV4cCI6MTY3MzUzMTY1N30.XNyNdGwe8xYS7RvPM-LcZph_ade8EfEV2fMV7WPSmZc",
}
var header []byte

func init() {
	header, _ = json.Marshal(m1)
}

func TestDeviceList(t *testing.T) {
	rep, err := helper.HttpGet(adminServiceAddr+"/device/list?page=1&size=20&name=", header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestDeviceCreate(t *testing.T) {
	m2 := map[string]string{
		"name":             "name",
		"product_identity": "1",
	}
	data, _ := json.Marshal(m2)

	rep, err := helper.HttpPost(adminServiceAddr+"/device/create", data, header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestDeviceModify(t *testing.T) {
	m2 := map[string]string{
		"identity":         "1",
		"name":             "name1",
		"product_identity": "1",
	}
	data, _ := json.Marshal(m2)

	rep, err := helper.HttpPut(adminServiceAddr+"/device/modify", data, header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestDeviceDelete(t *testing.T) {
	m2 := map[string]string{
		"identity": "1a6ecbfb-b3e6-4459-86cd-b43c61a22fa2",
	}
	data, _ := json.Marshal(m2)

	rep, err := helper.HttpDelete(adminServiceAddr+"/device/delete", data, header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestProductList(t *testing.T) {
	rep, err := helper.HttpGet(adminServiceAddr+"/product/list?page=1&size=20&name=", header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestProductCreate(t *testing.T) {
	m2 := map[string]string{
		"name": "name",
		"desc": "desc",
	}
	data, _ := json.Marshal(m2)

	rep, err := helper.HttpPost(adminServiceAddr+"/product/create", data, header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestProductModify(t *testing.T) {
	m2 := map[string]string{
		"identity": "1",
		"name":     "name1",
		"desc":     "desc2",
	}
	data, _ := json.Marshal(m2)

	rep, err := helper.HttpPut(adminServiceAddr+"/product/modify", data, header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestProductDelete(t *testing.T) {
	m2 := map[string]string{
		"identity": "1",
	}
	data, _ := json.Marshal(m2)

	rep, err := helper.HttpDelete(adminServiceAddr+"/product/delete", data, header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
