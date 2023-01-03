package test

import (
	"encoding/json"
	"fmt"
	"gitee/getcharzp/iot-platform/helper"
	"testing"
)

var emqxAddr = "http://192.168.1.8:18083/api/v5"

type createAuthUserRequest struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

func TestCreateAuthUser(t *testing.T) {
	// {"user_id":"go","password":"123456","is_superuser":false}
	c := createAuthUserRequest{
		UserId:   "user-id1",
		Password: "123456",
	}
	data, _ := json.Marshal(c)
	rep, err := helper.HttpPost(emqxAddr+"/authentication/password_based%3Abuilt_in_database/users", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestDeleteAuthUser(t *testing.T) {
	clientID := "user-id"
	rep, err := helper.HttpDelete(emqxAddr+"/authentication/password_based%3Abuilt_in_database/users/"+clientID, []byte{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestGetAuthUserList(t *testing.T) {
	rep, err := helper.HttpGet(emqxAddr + "/authentication/password_based%3Abuilt_in_database/users?limit=20&page=1")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
