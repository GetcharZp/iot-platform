package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	zrpc.RpcClientConf
	Mqtt struct {
		Broker   string
		ClientID string
		Password string
	}
}
